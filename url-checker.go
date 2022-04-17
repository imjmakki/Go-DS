package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is Down!\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d \n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}
	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}
}
