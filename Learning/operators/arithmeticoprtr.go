package main

import "fmt"

func main() {
	var a, b string = "fat", "her"
	fmt.Println(a + b)
	c, d := 10, 20
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)
	fmt.Println(c % d)
	c++
	fmt.Println(c)
	d--
	fmt.Println(c)

}
