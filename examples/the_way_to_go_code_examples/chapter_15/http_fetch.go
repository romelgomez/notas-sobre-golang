// httpfetch.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com")
	CheckError(err)
	data, err := ioutil.ReadAll(res.Body)
	CheckError(err)
	fmt.Printf("Got: %q", string(data))
}

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Get: %v", err)
	}
}
