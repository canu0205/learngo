package urlcheck

import (
	"errors"
	"fmt"
	"net/http"
)

// URLStatus struct
type URLStatus struct {
	URL    string
	Status string
}

var errRequestFailed = errors.New("Request failed")

// HitURL checks a url
func HitURL(url string, c chan<- URLStatus) {
	fmt.Println("Checking:", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		c <- URLStatus{URL: url, Status: "FAILED"}
	}
	c <- URLStatus{URL: url, Status: res.Status}
}
