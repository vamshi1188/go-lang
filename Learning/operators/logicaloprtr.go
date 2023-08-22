package main

import "fmt"

func main() {
	var a, b int = 1, 2
	c, d := true, false
	fmt.Println((a > 100) && (b > 100))
	fmt.Println(c || d)
	fmt.Println(c != d)

}
