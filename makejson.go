package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	person := make(map[string]string)
	var name string
	var address string

	fmt.Println("Please enter the name: ")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	name = scanner1.Text()
	person["name"] = name

	fmt.Println("Please enter the address: ")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	address = scanner2.Text()
	person["address"] = address

	response, err := json.Marshal(person)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(response))
}
