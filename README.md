# Car Rental API

Simple **Car Rental REST API** built with **Golang** and **PostgreSQL**.

This project provides CRUD operations for:

- Customers
- Cars
- Bookings

---

# Tech Stack

- Golang
- Gin Framework
- PostgreSQL
- REST API

---

# ERD (Entity Relationship Diagram)

Entities used in this system:

![ERD](https://github.com/primamaulana/Car-Rental-API/blob/main/erd/erd.png)

---

# Database Setup

Import **car_rental.sql** file from **database** folder or<br><br>
Create database manually:

```sql
CREATE DATABASE car_rental;
```

Use database:

```sql
\c car_rental
```

Create tables:

```sql
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    nik VARCHAR(20),
    phone_number VARCHAR(20)
);

CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    stock INT,
    daily_rent NUMERIC
);

CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id),
    car_id INT REFERENCES cars(id),
    start_rent DATE,
    end_rent DATE,
    total_cost NUMERIC,
    finished BOOLEAN
);
```

---

# How to Run Project

### 1 Install Dependencies

```
go mod tidy
```

or

```
go get github.com/gin-gonic/gin
go get github.com/lib/pq
```

---

### 2 Configure Database

Edit connection in `db.go`

```go
connStr := "host=localhost port=5432 user=postgres password=postgres dbname=car_rental sslmode=disable"
```

Adjust according to your local PostgreSQL credentials.

---

### 3 Run Application

```
go run .
```

Server will run on:

```
http://localhost:8080
```

---

# API Endpoint List

## Customer API

### Get All Customers

```
GET /customers
```

### Create Customer

```
POST /customers
```

Request Body:

```json
{
  "name": "John Doe",
  "nik": "3372093912739",
  "phone_number": "08123456789"
}
```

### Update Customer

```
PUT /customers/:id
```

### Delete Customer

```
DELETE /customers/:id
```

---

# Cars API

### Get All Cars

```
GET /cars
```

### Create Car

```
POST /cars
```

Request Body:

```json
{
  "name": "Toyota Camry",
  "stock": 2,
  "daily_rent": 500000
}
```

### Update Car

```
PUT /cars/:id
```

### Delete Car

```
DELETE /cars/:id
```

---

# Booking API

### Get All Bookings

```
GET /bookings
```

### Create Booking

```
POST /bookings
```

Request Body:

```json
{
  "customer_id": 1,
  "car_id": 1,
  "start_rent": "2026-03-16",
  "end_rent": "2026-03-18",
  "total_cost": 1000000,
  "finished": false
}
```

### Update Booking

```
PUT /bookings/:id
```

### Delete Booking

```
DELETE /bookings/:id
```

---

# Example API Flow

1 Create Customer

```
POST /customers
```

2 Create Car

```
POST /cars
```

3 Create Booking

```
POST /bookings
```

---
# Postman Collection

You can test all API endpoints using the Postman collection provided.

Import the file below into Postman:

**postman/car-rental-api.postman_collection.json**

---
# Author

Prima Maulana Hanan
