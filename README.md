# 🛠️ GeoMart - Backend

This is the backend service for the **GeoMart** e-commerce platform, built using **Go (Gin framework)** and **GORM** for ORM. It supports user authentication integration with Auth0, user profile management, order tracking, and serves as the API layer for the GeoMart frontend.

## ⚙️ Tech Stack

* **Go (Golang)** with Gin web framework
* **GORM** for ORM and MySQL/PostgreSQL database support
* **REST API** architecture
* **CORS** and **JSON** middleware
* **Auth0** integration for secure authentication
* **PayPal**-linked order endpoints (via frontend coordination)

## 📁 Project Structure

```
GeoMart-Backend/
├── controllers/ # Business logic and API handlers
├── models/ # GORM models for DB tables
├── routes/ # Route definitions and grouping
├── utils/ # Helper functions (e.g., database connection)
├── main.go # Entry point of the application
└── go.mod / go.sum # Go module dependencies
```

## 🚀 Features

* ✅ **User Creation/Check**: Create or validate user data on login
* 📄 **User Profile**: Fetch and update profile info (phone, address, etc.)
* 📦 **Order Handling**: Save orders received from frontend checkout
* 📡 **CORS-enabled APIs** for frontend integration
* 🔒 **Secured with Auth0 sub identifier**

## 🛠️ Setup Instructions

### 🔧 Prerequisites

* Go 1.20+ installed
* MySQL or PostgreSQL setup
* Auth0 application set up
* CORS enabled on frontend server (e.g. http://localhost:3000)

### 📦 Installation

```bash
git clone https://github.com/Nitinnannapaneni20/GeoMart-Backend.git
cd GeoMart-Backend
go mod tidy
```

### ⚙️ Environment Variables

Create a `.env` file in the root directory:

```env
DB_USER=your-db-username
DB_PASS=your-db-password
DB_NAME=geodb
DB_HOST=localhost
DB_PORT=3306
```

Make sure your database is running and matches the credentials.

### ▶️ Run the App

```bash
go run main.go
```

App will be live at: http://localhost:8080

## 🤝 Contributing

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature`
3. Commit your changes: `git commit -m "Add your feature"`
4. Push to GitHub: `git push origin feature/your-feature`
5. Open a pull request 🚀


By [@Nitinnannapaneni20](https://github.com/Nitinnannapaneni20)