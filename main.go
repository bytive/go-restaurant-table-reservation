package main

import (
	"fmt"
	"reservation/src/reservation"
)

func main() {
	fmt.Println("hi")
	reservation.CreateReservationTables()
	reservation.PrintReservationTables()
	reservation.TakeBookingInput()
}
