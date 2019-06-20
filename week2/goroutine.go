package main

import (
	"net/http"
	"fmt"
	"sync"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go healthcheck("https://github.com", &waitgroup)
	waitgroup.Add(1)
	go healthcheck("https://google.com", &waitgroup)
	waitgroup.Add(1)
	go healthcheck("https://pornhub.com", &waitgroup)
	waitgroup.Add(1)
	go healthcheck("https://thiendia.com", &waitgroup)
	waitgroup.Add(1)
	go healthcheck("https://zing.vn", &waitgroup)
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
