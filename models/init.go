package models

import (
	"time"
)

type Apartments struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	RoomCount          uint   `json:"room_count"`
	SleepingPlaceCount uint   `json:"sleeping_place_count"`
}

type Bookings struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	PhoneNumber   string    `json:"description"`
	ArrivalTime   time.Time `json:"arrival_time"`
	DepartureTime time.Time `json:"departure_time"`
	ApartmentID   uint      `json:"apartment_id" gorm:"foreignKey:ApartmentsID;references:ID"`
	Status        string    `json:"status"`
}

type Pictures struct {
	ID          uint   `json:"id"`
	ApartmentID uint   `json:"apartment_id" gorm:"foreignKey:ApartmentsID;references:ID"`
	FileName    string `json:"file_name"`
}
