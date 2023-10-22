package reservation

import (
	"fmt"
	"log"
	"time"
)

var bookings []ReservationBooking

func TakeBookingInput() {
	fmt.Println("Please enter (Y) to start inputting booking details")
	var userInput string
	_, err := fmt.Scanf("%s", &userInput)
	if err != nil {
		log.Fatal(err)
	}

	if userInput == "Y" {
		fmt.Println("Enter the no. of Guests")
		var noOfGuest int
		_, err := fmt.Scanf("%d", &noOfGuest)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("NoOfGuest", noOfGuest)
		isValid := ValidateNoOfGuest(noOfGuest)
		if isValid {
			HandleTableBooking(noOfGuest)
		} else {
			fmt.Println("Oops!! we don't have booking space available !!")
			TakeBookingInput()
		}

	} else {
		fmt.Println("User Input is", userInput)
		TakeBookingInput()
	}

}

func ValidateNoOfGuest(count int) bool {
	totalFree := 0
	for _, table := range tables {
		if table.IsEmpty {
			totalFree += table.NoOfSittings
			//
			// totalFree = totalFree + table.NoOfSittings
		}
	}
	return totalFree >= count
	// similar to ...
	// if totalFree>=count{
	// 	return true
	// }else{
	// 	return false
	// }
}

func HandleValidateTimeSlot() {

}

func HandleTableBooking(guest int) {
	// take user input for handling the booking
	TakeBookingsDetailsUserInput(guest)
	// validate the timeslot availablity
	// confirm the booking
	// move table to the same function
	TakeBookingInput()
}

func TakeBookingsDetailsUserInput(guests int) {
	var reservation ReservationBooking

	fmt.Print("Enter your name: ")
	fmt.Scanln(&reservation.PrimaryGuestName.Name)

	fmt.Print("Enter your email: ")
	fmt.Scanln(&reservation.PrimaryGuestName.Email)

	fmt.Print("Enter the date (e.g., DD/MM/YYYY): ")
	fmt.Scanln(&reservation.Date)

	fmt.Print("Enter the time (e.g., HH:MM AM/PM): ")
	fmt.Scanln(&reservation.Time)
	fmt.Println("TIME IOS", reservation.Time)

	// handle the guestNo.
	reservation.NoOfGuest = guests

	// Convert the reservation date and time to a time.Time object
	dateTimeFormat := "02/01/2006 03:04 PM"
	parsedDateTime, err := time.Parse(dateTimeFormat, reservation.Date+" "+reservation.Time)
	fmt.Println("errr", err)
	if err != nil {
		fmt.Println("Invalid date or time format. Reservation not added.")
		return
	}

	// Check for conflicts with existing reservations
	for _, existing := range bookings {
		existingDateTime, err := time.Parse(dateTimeFormat, existing.Date+" "+existing.Time)
		if err != nil {
			continue // Skip invalid date/time formats in existing reservations
		}

		if parsedDateTime.Equal(existingDateTime) {
			fmt.Println("A reservation already exists for this date and time. Reservation not added.")
			return
		}

		bookings = append(bookings, reservation)
		fmt.Println("Reservation added successfully!")
	}

}
