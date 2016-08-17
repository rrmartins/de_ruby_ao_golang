package main

import "fmt"

func foobar(str string) {
	fmt.Println("hello world,", str)
}

func foo() {
	fmt.Println("hello world")
}

func bar(str string) string {
	return fmt.Sprintf("hello world, %s", str)
}

func main() {
	foo()
	s := "Rodrigo Martins"
	foobar(s)
	str := bar(s)
	fmt.Println(str)
}
