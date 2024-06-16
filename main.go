package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func hitUrl(url string) error {
	fmt.Println("check :", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

func main() {
	// initialize
	// if you want to initialize an empty map, do like this.
	// if not your map goes nil and you can't write into nil.
	var results = make(map[string]string)

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
		result := "OK"
		err := hitUrl(url)
		if err != nil {
			result = "Failed"
		}
		results[url] = result
	}

	// key, val
	for url, result := range results {
		fmt.Println(url, result)
	}
}
