package main

import "fmt"

func main() {
	confernceName := "Go Confirence"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", confernceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var userTickets int
	//ask user for their name
	fmt.Println("Ismingiz nima?")
	fmt.Scan(&firstName)
	fmt.Println("Familyangiz nima?")
	fmt.Scan(&lastName)
	fmt.Println("Qancah ticket olmoqchisiz?")
	fmt.Scan(&userTickets)

	fmt.Printf("Rahmat %v %v , bizdan %v dona ticket harid qilganingiz uchun\n", firstName, lastName, userTickets)

}
