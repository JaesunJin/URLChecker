package main

import (
	"errors"
	"fmt"
	"net/http"
)

type urlCheckResult struct {
	url, result string
}

var errRequestFailed = errors.New("Request failed")

func main() {

	c := make(chan urlCheckResult)
	var resultArray []urlCheckResult

	urls := []string{
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.naver.com",
		"https://www.airbnb.com",
		"https://www.yahoo.co.jp",
		"https://www.dcinside.com",
		"https://www.daum.net",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://sc054117.synology.me",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		resultArray = append(resultArray, <-c)
	}
	for _, result := range resultArray {
		fmt.Println("URL : ", result.url, "RESULT : ", result.result)
	}
}
func hitURL(url string, c chan<- urlCheckResult) {
	// fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- urlCheckResult{url: url, result: "NG"}
	}
	c <- urlCheckResult{url: url, result: "OK"}
}
