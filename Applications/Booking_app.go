package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v  are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend ")

	for {

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

		isvalidName := len(firstName) >= 2 && len(lastName) >= 0
		isvalidEmail := strings.Contains(email, "@")
		isvalidUsertickets := userTickets > 0 && userTickets <= remainingTickets

		if isvalidEmail && isvalidName && isvalidUsertickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for bokking %v tickets. you will get a confirmation mail on %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v ticketes are remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("the  first names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println(" Our conference tickets are sold. come back next year.")
				break

			}
		} else {
			fmt.Println("Your input data is invalid. try again")

		}

	}

}
