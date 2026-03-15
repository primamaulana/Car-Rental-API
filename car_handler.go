package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {

	rows, _ := DB.Query("SELECT id,name,stock,daily_rent FROM cars")

	var cars []Car

	for rows.Next() {

		var car Car

		rows.Scan(&car.ID, &car.Name, &car.Stock, &car.DailyRent)

		cars = append(cars, car)
	}

	c.JSON(http.StatusOK, cars)
}

func CreateCar(c *gin.Context) {

	var car Car

	c.BindJSON(&car)

	err := DB.QueryRow(
		"INSERT INTO cars(name,stock,daily_rent) VALUES($1,$2,$3) RETURNING id",
		car.Name,
		car.Stock,
		car.DailyRent,
	).Scan(&car.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, car)
}

func UpdateCar(c *gin.Context) {

	id := c.Param("id")

	var car Car
	c.BindJSON(&car)

	_, err := DB.Exec(
		"UPDATE cars SET name=$1, stock=$2, daily_rent=$3 WHERE id=$4",
		car.Name,
		car.Stock,
		car.DailyRent,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteCar(c *gin.Context) {

	id := c.Param("id")

	DB.Exec("DELETE FROM cars WHERE id=$1", id)

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}