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
	status := "OK"

	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
		c <- URLStatus{URL: url, Status: status}
	}
	c <- URLStatus{URL: url, Status: status}
}
