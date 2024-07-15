package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	company  string
	location string
	workexp  string
	skill    string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=devops"
var idURL string = "https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx="

func main() {
	var jobs []extractedJob
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)

		break // 차단당하지 않기 위해
	}
	// fmt.Println(jobs)
	writeJobs(jobs)
	fmt.Println("Done, extracted ", len(jobs))

}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Company", "Location", "Workexp", "Skill"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobsSlice := []string{idURL + job.id, job.title, job.company, job.location, job.workexp, job.skill}
		jwErr := w.Write(jobsSlice)
		checkErr(jwErr)
	}

}

func getPage(page int) []extractedJob {
	var jobs []extractedJob

	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page+1)
	fmt.Println(pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close() // 리소스를 해제하고 연결을 반환

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("value")
	title, _ := card.Find(".area_job>h2>a").Attr("title")
	company := cleanStrings(card.Find(".area_corp>.corp_name>a").Text())
	location := cleanStrings(card.Find(".area_job>.job_condition>span>a").Text())

	var workexp strings.Builder
	card.Find(".area_job>.job_condition>span").Eq(1).Each(func(i int, s *goquery.Selection) {
		workexp.WriteString(s.Text())
	})

	var skill strings.Builder
	card.Find(".area_job>.job_sector>a").Each(func(j int, sel *goquery.Selection) {
		skill.WriteString(sel.Text())
		skill.WriteString(" ")
	})

	// fmt.Println(id, title, company, location, workexp.String(), skill.String())

	return extractedJob{id: id,
		title:    title,
		company:  company,
		location: location,
		workexp:  workexp.String(),
		skill:    skill.String()}

}

func getPages() int {
	pages := 0

	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close() // 리소스를 해제하고 연결을 반환

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err) // 로그 메시지를 출력하고 프로그램을 종료하는 데 사용
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

func cleanStrings(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
