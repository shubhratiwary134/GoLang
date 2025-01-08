package main

import "fmt"

func main() {
	var nameOfConference string = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("welcome this is %v \n", nameOfConference)
	fmt.Printf("Welcome to the booking app, the tickets available are %v and the total tickets are %v \n", remainingTickets, conferenceTickets)
	fmt.Printf("remaining ticket is of %T type \n", remainingTickets)
	var username string
	var userTickets int
	username = "tom"
	userTickets = 2
	fmt.Printf("user named %v booked %v tickets \n", username, userTickets)
}
