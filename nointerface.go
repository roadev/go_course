package main

import "fmt"

type dog struct{}

type fish struct{}

type bird struct{}

func (dog) walk() string {
	return "I'm a dog and I walk"
}

func (fish) swim() string {
	return "I'm a fish and I swim"
}

func (bird) fly() string {
	return "I'm a bird and I'm flying"
}

func moveDog(p dog) {
	fmt.Println(p.walk())
}

func moveFish(f fish) {
	fmt.Println(f.swim())
}

func moveBird(b bird) {
	fmt.Println(b.fly())
}

func main() {
	p := dog{}
	moveDog(p)
	f := fish{}
	moveFish(f)
	b := bird{}
	moveBird(b)

}