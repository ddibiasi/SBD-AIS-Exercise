# Exercise 2 — REST Server in Go (Chi Framework)

This exercise implements a simple REST API server in Go using the **Chi** router.  
The service exposes multiple endpoints for managing drink orders and returns JSON responses.

---

## 🧩 Implemented Endpoints

| Method | Endpoint | Description |
|:-------|:----------|:-------------|
| **GET** | `/api/menu` | Returns a list of available drinks. |
| **POST** | `/api/order` | Accepts a single order in JSON and stores it in memory. |
| **GET** | `/api/order/all` | Returns a list of all orders. |
| **GET** | `/api/order/totalled` | Returns a map of drink IDs and their total quantities ordered. |
| **GET** | `/` | Displays a simple dashboard summarizing orders. |

---

## How to Run the Server

Open a terminal (PowerShell recommended) and run:

```powershell
cd Exc_2/skeleton
go run .

If everything compiles correctly, the server will start on:

http://localhost:3000

How to Test the Endpoints

You can test your API using PowerShell commands:
# Add an order
Invoke-RestMethod -Method POST -Uri "http://localhost:3000/api/order" `
  -ContentType "application/json" `
  -Body '{"drink_id": 1, "amount": 2}'

# List all orders
Invoke-RestMethod -Uri "http://localhost:3000/api/order/all"

# Show total drinks ordered
Invoke-RestMethod -Uri "http://localhost:3000/api/order/totalled"

# Open dashboard in browser
start http://localhost:3000/

Notes

The data is stored in-memory, so it resets whenever the server restarts.

Models (Drink, Order) are defined in the model/ folder.

The repository logic is inside repository/db.go.

REST routes are implemented in rest/api.go.

The server entry point is main.go.

API documentation can be generated via:

./build-openapi-docs.sh

Author

Mustapha Oluwatoyin Gali
Exercise 2 — Software Architecture for Big Data (FH Upper Austria)


