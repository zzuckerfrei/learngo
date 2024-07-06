package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

// c chan<- resultResult : define direction from/to channel
func hitUrl(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}

func main() {
	// initialize
	// if you want to initialize an empty map, do like this.
	// if not your map goes nil and you can't write into nil.
	var results = make(map[string]string)

	// channel
	// goroutine과 메인함수 사이에 정보를 전달하기 위한 방법
	// 또는 goroutine 과 goroutine 간의 커뮤니케이션도 가능
	c := make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	// idx, val
	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, "is", status)
	}

}
