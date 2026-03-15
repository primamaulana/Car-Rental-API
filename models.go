package main

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	NIK         string `json:"nik"`
	PhoneNumber string `json:"phone_number"`
}

type Car struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	DailyRent float64 `json:"daily_rent"`
}

type Booking struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	CarID      int     `json:"car_id"`
	StartRent  string  `json:"start_rent"`
	EndRent    string  `json:"end_rent"`
	TotalCost  float64 `json:"total_cost"`
	Finished   bool    `json:"finished"`
}