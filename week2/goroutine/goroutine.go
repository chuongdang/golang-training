package main

import (
	"net/http"
	"fmt"
	"sync"
	"runtime"
)

func main() {
	urlList := []string{
		"https://github.com",
		"https://google.com",
		"https://pornhub.com",
		"https://thiendia.com",
		"https://zing.vn",
	}

	runtime.GOMAXPROCS(8)
	var waitgroup sync.WaitGroup
	waitgroup.Add(len(urlList))
	for _, url := range urlList {
		go healthcheck(url, &waitgroup)
	}
	waitgroup.Wait()
}

func healthcheck(url string, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Could not connect to", url)
	}
	fmt.Println(url, res.StatusCode)
}
