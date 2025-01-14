package main

import (
	"fmt"
	"strings"
)

func main() {
	nameOfConference := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string
	fmt.Printf("welcome this is %v \n", nameOfConference)
	greetUsers()
	fmt.Printf("Welcome to the booking app, the tickets available are %v and the total tickets are %v \n", remainingTickets, conferenceTickets)
	//Array
	//var bookings [50]string

	for {
		//userTicket logic
		var username string
		var userTickets uint
		var userEmail string
		fmt.Printf("enter your name \n")
		fmt.Scan(&username) // for taking the input you provide the pointer for the variable so that the scan function where to store the value inputted
		fmt.Printf("enter the tickets \n")
		fmt.Scan(&userTickets)
		fmt.Printf("enter the email of the user \n")
		fmt.Scan(&userEmail)
		isValidUserTicket := userTickets > 0 && userTickets <= remainingTickets
		isValidEmail := strings.Contains(userEmail, "@")
		if isValidUserTicket && isValidEmail {
			bookings = append(bookings, username)
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("user named %v booked %v tickets , you will be sent confirmation of the ticket at %v \n", username, userTickets, userEmail)
			fmt.Printf("remaining tickets are %v \n", remainingTickets)

			var noTicketRemaining bool = remainingTickets == 0

			if noTicketRemaining {
				fmt.Print("All tickets are booked")
				break
			}

		} else {
			if !isValidEmail {
				fmt.Print("wrong input email")
			}
			if !isValidUserTicket {
				fmt.Print("Invalid User Ticket")
			}
		}

	}

}

func greetUsers() {
	fmt.Print("welcome to our conference\n")
}
