package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int
	Body       string
}

func main() {

	urls := []string{"https://google.com", "https://snapp.ir"}
	responseChan := make(chan Response, len(urls))
	defer close(responseChan)
	for _, url := range urls {
		go Crawl(url, responseChan)
	}
	for response := range responseChan {
		time.Sleep(time.Second * 1)
		log.Println(response.Body)
	}
}

func Crawl(url string, ch chan Response) {
	client := http.Client{}
	resp, err := client.Get(url)

	if err == nil && resp.StatusCode == 200 {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panic("a problem detected in decoding")

		}
		ch <- Response{StatusCode: resp.StatusCode, Body: string(data)}
		resp.Body.Close()
	} else {
		log.Panic("a problem detected in connection")
	}
}
