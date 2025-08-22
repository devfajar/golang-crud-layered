# 📚 Golang REST API with PostgreSQL – Standalone Standard Library

A simple and clean RESTful API built using **Go (Golang)** standard library only – no external frameworks.
 This project showcases how to create a scalable, production-ready REST API using native Go, PostgreSQL.

---

## 🚀 Features

- ✅ Native HTTP server using `net/http`
- 🧱 Clean architecture (Handler → Service → Repository)
- 🗃️ PostgreSQL with `database/sql`
- 🔐 Environment config with `.env` & `godotenv`
- 🧪 Request validation using struct tags
- 🧰 Centralized helper utilities
- 📦 Fully modular & extensible project structure

---

## 📦 Tech Stack

| Layer        | Tool / Package                   |
|--------------|----------------------------------|
| Language     | [Go (Golang)](https://golang.org) |
| Database     | PostgreSQL                       |
| DB Driver    | [`lib/pq`](https://github.com/lib/pq) |
| Env Loader   | [`joho/godotenv`](https://github.com/joho/godotenv) |
| Logger       | [`rs/zerolog`](https://github.com/rs/zerolog) |

## 📁 Project Structure
├── config/ // DB config
├── data // DTO atau struct request and response
├── handler/ // HTTP request handlers
├── helper/ // Utilities: error handling, date parsing, etc.
├── model/ // Domain models
├── repository/ // Data access layer (SQL)
├── route/ // HTTP route definitions
├── service/ // Business logic layer
├── sql // query sql
├── .env // Environment variables
└── main.go // Entry point

## ⚙️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/devfajar/golang-crud-layered.git
cd golang-crud-layered
```

### 2. Create your .env file
```bash
cp .env.example .env
```
# Then edit .env and fill in your PostgreSQL credentials
DB_HOST=localhost

DB_PORT=5432

DB_USER=postgres

DB_PASSWORD=yourpassword

DB_NAME=yourtable

## 3. Run the application
```bash
go run main.go
```
