package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookings(c *gin.Context) {

	rows, _ := DB.Query("SELECT id,customer_id,car_id,start_rent,end_rent,total_cost,finished FROM bookings")

	var bookings []Booking

	for rows.Next() {

		var b Booking

		rows.Scan(&b.ID, &b.CustomerID, &b.CarID, &b.StartRent, &b.EndRent, &b.TotalCost, &b.Finished)

		bookings = append(bookings, b)
	}

	c.JSON(http.StatusOK, bookings)
}

func CreateBooking(c *gin.Context) {

	var b Booking

	c.BindJSON(&b)

	err := DB.QueryRow(
		"INSERT INTO bookings(customer_id,car_id,start_rent,end_rent,total_cost,finished) VALUES($1,$2,$3,$4,$5,$6) RETURNING id",
		b.CustomerID,
		b.CarID,
		b.StartRent,
		b.EndRent,
		b.TotalCost,
		b.Finished,
	).Scan(&b.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, b)
}

func UpdateBooking(c *gin.Context) {

	id := c.Param("id")

	var b Booking
	c.BindJSON(&b)

	_, err := DB.Exec(
		"UPDATE bookings SET customer_id=$1, car_id=$2, start_rent=$3, end_rent=$4, total_cost=$5, finished=$6 WHERE id=$7",
		b.CustomerID,
		b.CarID,
		b.StartRent,
		b.EndRent,
		b.TotalCost,
		b.Finished,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteBooking(c *gin.Context) {

	id := c.Param("id")

	DB.Exec("DELETE FROM bookings WHERE id=$1", id)

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}