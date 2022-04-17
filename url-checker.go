package main

import (
	"fmt"
	"net/http"
)

func checkAndSaveBody(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is Down!\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d \n", url, resp.StatusCode)
	}
}

func main() {

}
