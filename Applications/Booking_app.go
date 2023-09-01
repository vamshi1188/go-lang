package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings []string

func main() {

	greetuser()

	for {

		firstName, lastName, email, userTickets := getuserInput()

		isvalidName, isvalidEmail, isvalidUsertickets := validateUserinput(firstName, lastName, email, userTickets)

		if isvalidEmail && isvalidName && isvalidUsertickets {
			bookTickets(userTickets, firstName, lastName, email)
			firstNames := getFirstnames()
			fmt.Printf("the  first names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println(" Our conference tickets are sold. come back next year.")
				break

			}
		} else {
			if !isvalidName {
				fmt.Println("first name or lastname you entered is too short")
			}
			if !isvalidEmail {
				fmt.Println("the email does not contain '@' symbol")
			}
			if !isvalidUsertickets {
				fmt.Println("the number of tickets you entered is invalid")
			}

		}

	}
}
func greetuser() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v  are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend ")
}
func getFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}
func validateUserinput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	isvalidName := len(firstName) >= 2 && len(lastName) >= 0
	isvalidEmail := strings.Contains(email, "@")
	isvalidUsertickets := userTickets > 0 && userTickets <= remainingTickets
	return isvalidName, isvalidEmail, isvalidUsertickets
}
func getuserInput() (string, string, string, uint) {

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
	return firstName, lastName, email, userTickets
}
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for bokking %v tickets. you will get a confirmation mail on %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v ticketes are remaining for %v\n", remainingTickets, conferenceName)

}

// bouhcueqhcJHFIWJB
