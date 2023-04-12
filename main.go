package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int=50
	var remaningTickets uint=50
	
	fmt.Printf("Welcome to %v booking appliction\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n",conferenceTickets,remaningTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their name
	fmt.Print("Enter the user first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter the user last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter the user email: ")
	fmt.Scan(&email)
	fmt.Print("How much ticket you want to book: ")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v %v booked %v tickets using this email ID:%v\n",firstName,lastName,userTickets,email)
}  