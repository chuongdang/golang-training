package main

import (
	"net/http"
	"fmt"
)

func main() {
	healthcheck("https://github.com")
	healthcheck("https://google.com")
	healthcheck("https://pornhub.com")
	healthcheck("https://thiendia.com")
	healthcheck("https://zing.vn")
}

func healthcheck(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Could not connect to", url)
	}
	fmt.Println(url, res.StatusCode)
}
