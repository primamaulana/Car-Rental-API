package main

import "github.com/gin-gonic/gin"

func main() {

	InitDB()

	r := gin.Default()

	// Customer
	r.GET("/customers", GetCustomers)
	r.POST("/customers", CreateCustomer)
	r.PUT("/customers/:id", UpdateCustomer)
	r.DELETE("/customers/:id", DeleteCustomer)

	// Cars
	r.GET("/cars", GetCars)
	r.POST("/cars", CreateCar)
	r.PUT("/cars/:id", UpdateCar)
	r.DELETE("/cars/:id", DeleteCar)

	// Bookings
	r.GET("/bookings", GetBookings)
	r.POST("/bookings", CreateBooking)
	r.PUT("/bookings/:id", UpdateBooking)
	r.DELETE("/bookings/:id", DeleteBooking)

	r.Run(":8080")
}