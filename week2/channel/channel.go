package main

import (
	"net/http"
	"fmt"
)

func main() {
	urlList := []string{
		"https://github.com",
		"https://google.com",
		"https://pornhub.com",
		"https://thiendia.com",
		"https://zing.vn",
	}

	channel := make(chan string, len(urlList))
	for _, url := range urlList {
		go healthcheck(url, &channel)
	}
	var a string
	for _ = range urlList {
		a = <-channel
		fmt.Println(a)
	}
}

func healthcheck(url string, channel *chan string) {
	fmt.Println(channel)
	defer markDone(channel)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Could not connect to", url)
	}
	fmt.Println(url, res.StatusCode)
}

func markDone(channel *chan string) {
	*channel <- "Done"
}
