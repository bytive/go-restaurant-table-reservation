package reservation

type ReservationTable struct {
	TableNo      int  `json:"table_no"`
	NoOfSittings int  `json:"no_of_sittings"`
	IsEmpty      bool `json:"is_empty"`
}

type ReservationBooking struct {
	NoOfGuest           int            `json:"no_of_guest"`
	Date                string         `json:"date"`
	Time                string         `json:"slot_time"`
	PrimaryGuestName    BookingGuest   `json:"primary_guest_name"`
	SecondaryGuestsName []BookingGuest `json:"secondary_guest_name"`
}

type BookingGuest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
