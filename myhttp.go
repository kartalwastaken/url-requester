package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

const MaximumParallelRequestCount int = 10

func main() {
	parallel := flag.Int(
		"parallel", MaximumParallelRequestCount,
		"The parallel request count.",
	)
	flag.Parse()
	if *parallel == 0 {
		fmt.Println("Parallel request count can't be zero.")
		return
	}

	listOfURLs := flag.Args()
	if len(listOfURLs) == 0 {
		fmt.Println("At least one url must be defined for the request.")
		return
	}

	urlChan := make(chan string)
	wg := new(sync.WaitGroup)

	// Adding routines to workgroup and running then
	for i := 0; i < *parallel; i++ {
		wg.Add(1)
		go sendRequest(urlChan, wg)
	}

	// Processing all urls by spreading them to `free` goroutines
	for _, url := range listOfURLs {
		urlChan <- url
	}

	close(urlChan)
	wg.Wait()
}

func sendRequest(urlChan chan string, wg *sync.WaitGroup) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	for url := range urlChan {
		url := normalizeURL(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("error requesting url: %s, err: %s \n", url, err)
			return
		}

		hash := md5.New()
		err = resp.Write(hash)
		if err != nil {
			fmt.Printf("error hashing response. url: %s, response code: %d \n", url, resp.StatusCode)
			return
		}
		fmt.Printf("%s %s\n", url, hex.EncodeToString(hash.Sum(nil)))
	}
}

func normalizeURL(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Printf("error parse url: %s, err: %s \n", rawURL, err)
		return ""
	}
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	return u.String()
}
