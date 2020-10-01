package main

import "fmt"

func main() {
	map1 := make(map[string]int)

	map1["a"] = 8

	fmt.Println(map1)
	fmt.Println(map1["a"])
}