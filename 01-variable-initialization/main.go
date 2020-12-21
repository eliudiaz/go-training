package main

import "fmt"

func main() {
	var a string
	var b int

	a = "five"
	b = 5

	msg := ""
	msg = fmt.Sprintf("a = %s and b = %d", a, b)
	fmt.Println(msg)
}
