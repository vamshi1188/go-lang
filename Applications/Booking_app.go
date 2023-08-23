package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings [50]string
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v  are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend ")
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask for user
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter no of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("the whole array:%v\n", bookings)

	fmt.Printf("Thank you %v %v for bokking %v tickets. you will get a confirmation mail on %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v ticketes are remaining for %v\n", remainingTickets, conferenceName)
}
