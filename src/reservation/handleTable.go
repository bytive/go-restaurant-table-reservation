package reservation

import (
	"encoding/json"
	"fmt"
)

var tables []ReservationTable

func CreateReservationTables() []ReservationTable {
	table1 := ReservationTable{TableNo: 1, NoOfSittings: 5, IsEmpty: true}
	table2 := ReservationTable{TableNo: 2, NoOfSittings: 10, IsEmpty: true}
	table3 := ReservationTable{TableNo: 3, NoOfSittings: 3, IsEmpty: true}
	table4 := ReservationTable{TableNo: 4, NoOfSittings: 8, IsEmpty: true}

	tables = append(tables, table1, table2, table3, table4)

	return tables
}

func PrintReservationTables() {
	s, _ := json.MarshalIndent(tables, "", "\t")

	fmt.Println("reservation tables are : ", string(s))
}
