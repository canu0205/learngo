package main

import (
	"fmt"

	"github.com/canu0205/learngo/urlcheck"
)

func main() {
	results := make(map[string]string)
	c := make(chan urlcheck.URLStatus)

	urls := []string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.airbnb.com",
		"https://www.reddit.com",
	}

	for _, url := range urls {
		go urlcheck.HitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.URL] = result.Status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}
