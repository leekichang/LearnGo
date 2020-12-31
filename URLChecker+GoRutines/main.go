package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string) // var results = make(map[string]string) 랑 같음.
	c := make(chan result)
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
	results["hello"] = "Hello"
	for _, url := range urls {
		go hitURL(url, c)
	}
	time.Sleep(time.Second)
}

func hitURL(url string, c chan<- result) { //chan<- 하면 send only
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}
