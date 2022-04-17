package main

import (
	"fmt"
	"net/http"
)

func checkAndSaveBody(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

}
