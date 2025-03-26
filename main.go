package main

import "fmt"

func main() {
	 conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint= 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to Our %v Booking App\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are remaining\n",conferenceTickets, remainingTickets)
	fmt.Println("Book Your tickets Now!!")

	var bookings [50]string
	
	var firstName string 
	var lastName string 
	var email string 
	var userTickets uint
	fmt.Print("Enter Your First Name: ")
	fmt.Scan(&firstName)
	
	fmt.Print("Enter Your Last Name: ")
	fmt.Scan(&lastName)
	
	fmt.Print("Enter Your Email: ")
	fmt.Scan(&email)
	
	fmt.Print("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)
	
	bookings[0] = firstName+" "+lastName

	remainingTickets -= userTickets

	fmt.Printf("Whole bookings array: %v\n", bookings)
	fmt.Printf("First booking: %v\n", bookings[0])
	fmt.Printf("bookings type: %T\n", bookings)
	fmt.Printf("bookings length: %v\n", len(bookings))
	
	fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets for %v. Please check your email %v for more details\n", firstName, lastName, userTickets, conferenceName, email)

	fmt.Printf("We have %v tickets remaining\n", remainingTickets)
}
