package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	channel := make(chan string)


	servers := []string {
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
	}

	i := 0

	for {
		if i > 2 {
			break
		}
		for _, server := range servers {
			go checkServer(server, channel)
		}
		time.Sleep(1*time.Second)
		fmt.Println(<-channel)
		i++
	}
	

	executionTime := time.Since(start)

	fmt.Println("execution time %s\n", executionTime)

}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)

	if err != nil {
		// fmt.Println(server, " is down :(")
		channel <- server + " is not available"
	} else {
		// fmt.Println(server, " is running !")
		channel <- server + "it works! :D"
	}
}