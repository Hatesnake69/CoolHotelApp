package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

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

var dummyApartments = []Apartments{
	{ID: 1, Name: "Комната на первом этаже", Description: "Комната с видом на море", RoomCount: 1, SleepingPlaceCount: 2},
	{ID: 2, Name: "Комната на втором этаже", Description: "Комната с видом на море", RoomCount: 2, SleepingPlaceCount: 3},
	{ID: 3, Name: "Комната на третьем этаже", Description: "Комната с видом на море", RoomCount: 3, SleepingPlaceCount: 4},
}

var dummyBooking = []Bookings{
	{ID: 1, Name: "Вася", PhoneNumber: "+7(111)1111111", ArrivalTime: time.Now(), DepartureTime: time.Now(),
		ApartmentID: 1, Status: "free"},
}

var dummyPictures = []Pictures{
	{ID: 1, FileName: "109270328975140283.jpg", ApartmentID: 2},
	{ID: 2, FileName: "786343127693487123.jpg", ApartmentID: 1},
	{ID: 3, FileName: "109823409127364322.jpg", ApartmentID: 3},
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&Apartments{})
	db.AutoMigrate(&Bookings{})
	db.AutoMigrate(&Pictures{})

	if db.Find(&Apartments{}).RowsAffected == 0 &&
		db.Find(&Bookings{}).RowsAffected == 0 &&
		db.Find(&Bookings{}).RowsAffected == 0 {
		for _, item := range dummyApartments {
			db.Create(&item)
		}
		for _, item := range dummyPictures {
			db.Create(&item)
		}
		for _, item := range dummyBooking {
			db.Create(&item)
		}
	}

	router := gin.Default()
	router.GET("/apartments", getApartments)
	router.Run("localhost:8080")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func getApartments(c *gin.Context) {
	var result []Apartments
	db.Find(&result)
	c.IndentedJSON(http.StatusOK, dummyApartments)
}
