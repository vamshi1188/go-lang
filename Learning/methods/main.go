package main

import "fmt"

func main() {
	fmt.Println("welcome to methods in go lang")
	vamshi := User{"saivamshi", "saiva,m.email", true, 20}
	vamshi.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user online: ", u.Status)

}
