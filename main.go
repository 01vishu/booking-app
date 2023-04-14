package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remaningTickets uint = 50

	fmt.Printf("Welcome to %v booking appliction\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remaningTickets)
	fmt.Println("Get your tickets here to attend")

	var booking []string

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name
	fmt.Print("Enter the user first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter the user last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter the email address: ")
	fmt.Scan(&email)
	fmt.Print("How much ticket you want to book: ")
	fmt.Scan(&userTickets)

	remaningTickets = remaningTickets - userTickets
	// booking[0]=firstName +" " + lastName
	booking = append(booking, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will recevie a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remaningTickets, conferenceName)

	fmt.Printf("These are all our bookings %v\n", booking)
}
