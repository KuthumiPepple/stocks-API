# Stocks API

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This API provides endpoints for managing stocks

## Getting started
To run this application as-is:
1. Clone the repository and navigate to the project directory
2. Run `go mod tidy` command to download all the dependencies.
3. Create a PostgreSQL database. Create a table named `stocks`. The columns should align with the `Stock` model in the [models package](models/models.go) with the `stockid` column set as `SERIAL PRIMARY KEY`
4. Edit the .env file to use your database details.
5. Run `go run main.go` command to start the application

## Endpoints
| Endpoint | Method | Description |
| -------- | ------ | ----------- |
| /api/stocks   | POST   | Insert a new stock |
| /api/stocks   | GET    | Get all stocks |
| /api/stocks/{id} | GET    | Get stock by ID |
| /api/stocks/{id} | PUT    | Update an existing stock |
| /api/stocks/{id} | DELETE | Delete a stock |

## Consuming the API
Use any http client program to make HTTP requests to the running server. Here are some Request examples:
### **Insert a new Stock**
#### **Request**
POST /api/stocks
```json
{
    "symbol": "AB",
    "price": 19,
    "company": "Company1"
}
```
#### **Response**
201 Created
```json
{
    "id": 1,
    "message": "stock created successfully"
}
```
### **Get all Stocks**
#### **Request**
GET /api/stocks
#### **Response**
200 OK
```json
[
    {
        "stockid": 1,
        "symbol": "AB",
        "price": 19,
        "company": "Company1"
    },
    {
        "stockid": 2,
        "symbol": "CDE",
        "price": 152,
        "company": "Company2"
    },
    ...
    ...
]
```
### **Get Stock by ID**
#### **Request**
GET /api/stocks/1
#### **Response**
200 OK
```json
{
    "stockid": 1,
    "symbol": "AB",
    "price": 19,
    "company": "Company1"
}
```
### **Update an existing Stock**
#### **Request**
PUT /api/stocks/1
```json
{
    "price": 25
}
```
#### **Response**
200 OK
```json
{
    "id": 1,
    "message": "stock updated successfully. Total records affected: 1"
}
```
### **Delete a Stock**
#### **Request**
DELETE /api/stocks/2
#### **Response**
200 OK
```json
{
    "id": 2,
    "message": "Stock deleted. Total records affected: 1"
}
```