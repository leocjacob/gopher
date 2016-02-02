package main

import "fmt"

//untyped const
const name = "leo"
const num = 9999999
const num2 = 24

//typed const
const nme string = "hello"
const numb int = 43
const result int = num * num2

func main() {

	fmt.Print("product=", result)
}
