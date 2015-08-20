package main

import (
	"fmt"
	"net/http"
	"sync"
)

func check(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err == nil && resp.StatusCode == 200 {
		fmt.Println(url, "OK")
		return
	}
	fmt.Println(url, "NOK")
}

func main() {
	wg := &sync.WaitGroup{}
	urls := []string{
		"http://janus.hotelurbano.com/healthcheck",
		"http://novabusca.hotelurbano.com/healthcheck",
		"http://hotels.hotelurbano.com/healthcheck",
		"http://search.api.grupohu.com.br/api/healthcheck",
	}
	for _, url := range urls {
		wg.Add(1)
		go check(wg, url)
	}
	wg.Wait()
}
