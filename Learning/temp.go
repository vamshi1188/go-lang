package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// a := "sai"
	// b := "vamshi"
	// c := strings.Compare(a, b)
	// fmt.Println(c)

	//
	reader := bufio.NewReadWriter(os.Stdin)

	myname = reader.ReadString('\n')
	fmt.Println(myname)

}
