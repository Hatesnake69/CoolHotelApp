package main

import (
	"fmt"
	"net/http"
	"time"
)

type Apartment struct {
	ID                 uint
	Name               string
	Description        string
	RoomCount          uint
	SleepingPlaceCount uint
}

type Booking struct {
	ID            uint
	Name          uint
	PhoneNumber   string
	ArrivalTime   time.Time
	DepartureTime time.Time
	ApartmentID   Apartment
	Status        string
}

type Pictures struct {
	ID       uint
	FileName string
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
