package main

import "fmt"

func main(){
	conferenceName:= "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T, ramaingTicket is %T, conferenceTickets is %T \n",conferenceName,remainingTickets,conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n",conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	var userName string
	var userTickets int

	//ask  the user for their name
	fmt.Println("Enter your name ")
	fmt.Scan(&userName)


	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n",userName,userTickets)	
}