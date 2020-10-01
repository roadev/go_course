package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string
	barr := make([]byte, 1000)

	type name struct {
		fname string
		lname string
	}

	fmt.Println("Please enter the file name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename = scanner.Text()

	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	nb, err := f.Read(barr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(nb)

	namesString := string(barr)
	names := strings.Split(namesString, "\n")

	var sliceOfNames []name

	for i := 0; i < len(names); i++ {
		splittedName := strings.Split(names[i], " ")
		nameStruct := name{
			fname: splittedName[0],
			lname: splittedName[1],
		}

		sliceOfNames = append(sliceOfNames, nameStruct)
	}

	fmt.Println(names)
	fmt.Println("------------------")
	fmt.Println("Slice of name with two fields: ")

	for i := 0; i < len(sliceOfNames); i++ {
		fmt.Println(sliceOfNames[i])
		fmt.Println("first name:", sliceOfNames[i].fname, ", last name:", sliceOfNames[i].lname)
	}
}
