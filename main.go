package main

import (
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/zzuckerfrei/learngo/scrapper"
)

const FILE_NAME string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(FILE_NAME) // 사용자가 다운로드 한 후, 서버에서는 사용하지 않을 파일 삭제

	today := time.Now().Format("20060102")

	term := strings.ToLower(scrapper.CleanStrings(c.FormValue("term")))
	scrapper.Scrape(term)

	return c.Attachment(FILE_NAME, "saramin.com_"+term+"_"+today+"_"+FILE_NAME) // c.Attachment : 첨부파일 리턴
}

func main() {
	e := echo.New()

	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	e.Logger.Fatal(e.Start(":1323"))
}
