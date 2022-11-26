package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Apartments struct {
	ID                 uint
	Name               string
	Description        string
	RoomCount          uint
	SleepingPlaceCount uint
}

type Bookings struct {
	ID            uint
	Name          uint
	PhoneNumber   string
	ArrivalTime   time.Time
	DepartureTime time.Time
	ApartmentID   uint
	Status        string
}

type Pictures struct {
	ID       uint
	FileName string
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	db.Migrator().CreateTable(&Pictures{})
	db.Migrator().CreateTable(&Apartments{})
	db.Migrator().CreateTable(&Bookings{})
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
