package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to Our %v Booking App\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are remaining\n",conferenceTickets, remainingTickets)
	fmt.Println("Book Your tickets Now!!")

	var userName string
	userName = "John Doe"
	fmt.Println(userName)
}
