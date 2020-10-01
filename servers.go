package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	servers := []string {
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
	}

	for _, server := range servers {
		checkServer(server)
	}

	executionTime := time.Since(start)

	fmt.Println("execution time %s\n", executionTime)

}

func checkServer(server string) {
	_, err := http.Get(server)

	if err != nil {
		fmt.Println(server, " is down :(")
	} else {
		fmt.Println(server, " is running !")
	}
}