package main

import "fmt"

func main() {
	 conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int= 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to Our %v Booking App\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are remaining\n",conferenceTickets, remainingTickets)
	fmt.Println("Book Your tickets Now!!")

	var userName string 
	var userTickets int
	fmt.Print("Enter Your Name: ")
	fmt.Scan(&userName)

	fmt.Printf("User %v booked %v tickets", userName, userTickets)

}
