package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

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

// scrape Saramin by term
func Scrape(term string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + term
	var idURL string = "https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx="

	var jobs []extractedJob
	c := make(chan []extractedJob)
	var wg sync.WaitGroup

	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		wg.Add(1) // goroutine 카운트
		go getPage(i, baseURL, c)

		// break // 차단당하지 않기 위해 테스트용
	}

	// 모든 goroutine 완료되면 채널 닫기위한 goroutine 생성 -> 안정성
	go func() {
		wg.Wait() // 모든 goroutine이 종료되어 카운트가 일치할 때까지 대기
		close(c)
	}()

	// how many goroutine?
	for i := 0; i < totalPages; i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob...)
	}

	writeJobs(jobs, idURL)

	fmt.Println("Done, extracted ", len(jobs))
}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	var wg sync.WaitGroup

	pageURL := url + "&recruitPage=" + strconv.Itoa(page+1)
	fmt.Println(pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close() // 리소스를 해제하고 연결을 반환

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) {
		wg.Add(1)              // goroutine 카운트
		go extractJob(card, c) // create goroutine
	})

	// 모든 goroutine 완료되면 채널 닫기위한 goroutine 생성 -> 안정성
	go func() {
		wg.Wait() // 모든 goroutine이 종료되어 카운트가 일치할 때까지 대기
		close(c)
	}()

	// put a job into jobs
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

// send something to channel
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title, _ := card.Find(".area_job>h2>a").Attr("title")
	company := CleanStrings(card.Find(".area_corp>.corp_name>a").Text())
	location := CleanStrings(card.Find(".area_job>.job_condition>span>a").Text())

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

	c <- extractedJob{id: id,
		title:    title,
		company:  company,
		location: location,
		workexp:  workexp.String(),
		skill:    skill.String()}

}

func getPages(url string) int {
	pages := 0

	res, err := http.Get(url)
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

func writeJobs(jobs []extractedJob, idURL string) {
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

func CleanStrings(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
