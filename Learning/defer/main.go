package main

import "fmt"

func main() {
	fmt.Println("welcome to defers in go lang")
	mydefer()

}
func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
