package main

import (
	"net/http"
	"fmt"
	"io"
)

type webWriter struct{}

func (webWriter) Write (p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}


func main() {
	response, err := http.Get("https://juanroa.me")
	if err != nil {
		fmt.Println(err)
	}
	e := webWriter{}
	io.Copy(e, response.Body)
}