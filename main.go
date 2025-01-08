package main

import "fmt"

func main() {
	nameOfConference := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("welcome this is %v \n", nameOfConference)
	fmt.Printf("Welcome to the booking app, the tickets available are %v and the total tickets are %v \n", remainingTickets, conferenceTickets)
	fmt.Printf("remaining ticket is of %T type \n", remainingTickets)

	//userTicket logic
	var username string
	var userTickets int
	var userEmail string
	fmt.Printf("enter your name \n")
	fmt.Scan(&username)
	fmt.Printf("enter the tickets \n")
	fmt.Scan(&userTickets)
	fmt.Printf("enter the email of the user \n")
	fmt.Scan(&userEmail)

	fmt.Printf("user named %v booked %v tickets \n, you will be sent confirmation of the ticket at %v \n", username, userTickets, userEmail)
}
