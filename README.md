# Altacommerce API Project
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)
[![Go Reference](https://img.shields.io/badge/midtrans-reference-blue?logo=Midtrans&logoColor=white)](https://github.com/Midtrans/midtrans-go)

# Table of Content
- [Description](#description)
- [How to Use](#how-to-use)
- [Database Schema](#database-schema)
- [Testing Coverage]($testing-coverage)
- [Feature](#feature)
- [Endpoints](#endpoints)
- [Credits](#credits)

# Description
Altaecommerce merupakan project group kolaborasi BE & FE pada pembelajaran alterra academy, Repo ini hasil kontribusi Galang, Alka, Rizki sebagai tim BE8. Altaecommerce adalah Aplikasi Rest Server project E-Commerce
# Database Schema
![ERD](https://github.com/t3be8/altacommerce/blob/main/screenshoot/be8group.png)

# Testing Coverage
Implement Unit Testing average 


# Feature
List of overall feature in this Project (To get more details see the API Documentation below)
| No.| Feature        | Keterangan                                                             |
| :- | :------------- | :--------------------------------------------------------------------- |
| 1. | Register       | Authentication process                                                 |
| 2. | Login          | Authentication process                                                 |
| 3. | CRUD Product   | Create, Read, Update, and Delete Ingredient of the recipe in system    |
| 4. | CRUD Cart      | Add Product to Cart,                                                   |
| 5. | Order          | Create Order, Cancel Order,   Payout Order                             |


# How to Use
- Clone this repository in your $PATH:
```
$ git clone https://github.com/t3be8/altacommerce.git
```
- Cp file .env based on this project 
``
cp sample-env .env
``
- Don't forget to create database name as you want in your MySQL
- Run program with command
```
go run main.go
```
# Endpoints
Read the API documentation here [API Endpoint Documentation](https://app.swaggerhub.com/apis/altacommerce/alta-e_commerce_api/1.0.1) (Swagger)

# Credits
- [Galang Adi Puranto](https://github.com/adeeplearn) (Author)
- [Alka Prasetya](https://github.com/alkaprasetya) (Author)
- [Rizki Firdaus](https://github.com/marthadinatarf) (Author)

# Spesial Support
- [Jerry Young](https://github.com/jackthepanda96) (Mentor)
- [Pringgo GW]
- [Fahri]
