package main

import "fmt"

func main() {
	var i int
	i = 4
	fmt.Printf("The value of i is %d\n", i)

	i = "test with error"
	fmt.Printf("The new value of i is %s\n", i)
}
