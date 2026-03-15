package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {

	rows, err := DB.Query("SELECT id,name,nik,phone_number FROM customers")

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var customers []Customer

	for rows.Next() {

		var cust Customer

		rows.Scan(&cust.ID, &cust.Name, &cust.NIK, &cust.PhoneNumber)

		customers = append(customers, cust)
	}

	c.JSON(http.StatusOK, customers)
}

func CreateCustomer(c *gin.Context) {

	var cust Customer

	c.BindJSON(&cust)

	err := DB.QueryRow(
		"INSERT INTO customers(name,nik,phone_number) VALUES($1,$2,$3) RETURNING id",
		cust.Name,
		cust.NIK,
		cust.PhoneNumber,
	).Scan(&cust.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cust)
}

func UpdateCustomer(c *gin.Context) {

	id := c.Param("id")

	var cust Customer
	c.BindJSON(&cust)

	_, err := DB.Exec(
		"UPDATE customers SET name=$1, nik=$2, phone_number=$3 WHERE id=$4",
		cust.Name,
		cust.NIK,
		cust.PhoneNumber,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteCustomer(c *gin.Context) {

	id := c.Param("id")

	_, err := DB.Exec("DELETE FROM customers WHERE id=$1", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}