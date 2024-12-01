# E-Commerce Backend API

This project is a RESTful API built with Go and Gin for an e-commerce application. The API provides endpoints to manage users, products, cart functionality, addresses, and purchases.

---

## 🚀 Features

- **User Management**: Manage user accounts and authentication.
- **Product Management**: Retrieve and manage product data.
- **Cart Operations**: Add, remove, list items, and checkout from the cart.
- **Address Management**: Add, edit, and delete user addresses.
- **Instant Purchase**: Support for direct item purchases.

---

## 📋 Prerequisites

1. **Go**: Version 1.19 or later.
2. **Database**: MongoDB.
3. **Environment Variables**: Set up the necessary environment variables.

---

## 🛠 Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/Bitrox-Technology/Ecommerce-Backend.git
cd Ecommerce-Backend

.env

PORT=8000
MONGODB_URI=mongodb://localhost:27017
JWT_SECRET=your_jwt_secret

run server
go run main.go

