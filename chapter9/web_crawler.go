package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	urls := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.twitter.com",
		"http://www.cnn.com",
		"http://www.konga.com",
		"http://www.amazon.com",
		"http://www.aliexpress.com",
		"http://www.jumia.com",
	}

	homepageChan := make(chan HomepageUrl)

	for _, url := range urls {

		go func(url string) {
			res, err := http.Get(url)

			if err != nil {
				panic(err)
			}

			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)

			if err != nil {
				panic(err)
			}

			pageSize := len(bs)
			pageData := HomepageUrl{url, pageSize}

			fmt.Println(url, pageSize)

			homepageChan <- pageData

		}(url)
	}

	biggestSize := HomepageUrl{"", 0}

	for range urls {
		currentHomePageInChan := <-homepageChan
		if currentHomePageInChan.pageSize > biggestSize.pageSize {
			biggestSize = currentHomePageInChan
		}
	}

	fmt.Println("Url with biggest homepage size", biggestSize.url, ", pagesize:", biggestSize.pageSize)

}

type HomepageUrl struct {
	url      string
	pageSize int
}
