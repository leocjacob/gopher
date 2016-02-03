package main

import "fmt"

type counter int

func main() {
	var count counter
	fmt.Print("value of variable=", count)
	fmt.Printf("Type: %T", count)
}
