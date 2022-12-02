package main

import (
	"cool-hotel-app/db"
	"cool-hotel-app/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var dummyApartments = []models.Apartments{
	{ID: 1, Name: "Комната на первом этаже", Description: "Комната с видом на море", RoomCount: 1, SleepingPlaceCount: 2},
	{ID: 2, Name: "Комната на втором этаже", Description: "Комната с видом на море", RoomCount: 2, SleepingPlaceCount: 3},
	{ID: 3, Name: "Комната на третьем этаже", Description: "Комната с видом на море", RoomCount: 3, SleepingPlaceCount: 4},
}

var dummyBooking = []models.Bookings{
	{ID: 1, Name: "Вася", PhoneNumber: "+7(111)1111111", ArrivalTime: time.Now(), DepartureTime: time.Now(),
		ApartmentID: 1, Status: "free"},
}

var dummyPictures = []models.Pictures{
	{ID: 1, FileName: "109270328975140283.jpg", ApartmentID: 2},
	{ID: 2, FileName: "786343127693487123.jpg", ApartmentID: 1},
	{ID: 3, FileName: "109823409127364322.jpg", ApartmentID: 3},
}

func main() {
	db.PostgresDB.AutoMigrate(&models.Apartments{})
	db.PostgresDB.AutoMigrate(&models.Bookings{})
	db.PostgresDB.AutoMigrate(&models.Pictures{})

	if db.PostgresDB.Find(&models.Apartments{}).RowsAffected == 0 &&
		db.PostgresDB.Find(&models.Bookings{}).RowsAffected == 0 &&
		db.PostgresDB.Find(&models.Bookings{}).RowsAffected == 0 {
		for _, item := range dummyApartments {
			db.PostgresDB.Create(&item)
		}
		for _, item := range dummyPictures {
			db.PostgresDB.Create(&item)
		}
		for _, item := range dummyBooking {
			db.PostgresDB.Create(&item)
		}
	}

	router := gin.Default()
	router.GET("/apartments", getApartments)
	router.POST("/post-picture", postPicture)
	router.Run("localhost:8080")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func getApartments(c *gin.Context) {

	var result []models.Apartments
	db.PostgresDB.Find(&result)
	c.IndentedJSON(http.StatusOK, result)
}

func postPicture(c *gin.Context) {
	var newPicture models.Pictures
	if err := c.BindJSON(&newPicture); err != nil {
		fmt.Println(err)
		return
	}
	db.PostgresDB.Create(&newPicture)
	c.IndentedJSON(http.StatusCreated, newPicture)
}
