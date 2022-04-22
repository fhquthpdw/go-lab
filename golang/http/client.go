package main

import (
	"net/http"
	"sync"
	"time"
)

var _clientTimeout = 4 // this timeout should be less than server.WriteTimeout
var _httpClient = newHTTPClient(_clientTimeout)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func newHTTPClient(timeout int) HTTPClient {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(400)
	for i := 0; i < 400; i++ {
		go func() {
			defer wg.Done()

			req, _ := http.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
			_httpClient.Do(req)
		}()

	}
	wg.Wait()
}
