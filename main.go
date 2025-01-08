package main

import "fmt"

func main() {
	var nameOfConference = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets
	fmt.Printf("welcome this is %v \n", nameOfConference)
	fmt.Print("Welcome to the booking app, the tickets available are", remainingTickets)

}
