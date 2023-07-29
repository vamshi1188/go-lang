package main

import "fmt"

func main() {
	a, b := 10, 20
	a += b
	fmt.Println(a)
	a -= b
	fmt.Println(a)
	a *= b
	fmt.Println(a)
	a /= b
	fmt.Println(a)
	a %= b
	fmt.Println(a)

}
