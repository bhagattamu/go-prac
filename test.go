package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close() // Make sure we close the body
	ctype := res.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("Can't find content type header")
	}
	return ctype, nil
}

func main() {
	ctype, err := contentType("https://www.bhagatgurung.com.np")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
