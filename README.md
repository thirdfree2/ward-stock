# 🧼 Go Clean Architecture - WARD-STOCK-BACKEND

โปรเจกต์นี้ใช้แนวทาง **Clean Architecture** ด้วยภาษา Go (Golang) โดยมีการแยกแต่ละเลเยอร์อย่างชัดเจน เพื่อให้โค้ดสามารถทดสอบง่าย ขยายต่อได้ และแยก concern อย่างเป็นระบบ

---

## 🔧 Stack ที่ใช้

- 🐹 Go 1.20+
- 🧱 GORM (ORM สำหรับเชื่อมต่อ PostgreSQL)
- 🌐 Gin (Web Framework)
- 🐘 PostgreSQL
- 🐳 Docker + docker-compose
- 📁 Clean Architecture Principles

---

## 📦 Folder Structure

WARD-STOCK-BACKEND/
├── cmd/ # Entry point (main.go)
├── internal/
│ ├── domain/ # Entity + Repository Interface
│ ├── usecase/ # Business logic (interfaces & impl)
│ ├── delivery/
│ │ └── http/ # HTTP handlers (Gin)
│ └── infrastructure/
│ └── postgres/ # DB access layer
├── model/ # (legacy or transitional structs)
├── repository/ # (optional: interface if not in domain)
├── .env, .env.development # Environment config
├── docker-compose.yml # For local container setup
└── go.mod / go.sum

## 🧱 Clean Architecture Layer Diagram

[ delivery (handler/http) ] ← Gin HTTP API
                    ↓
[ usecase ] ← ธุรกิจหลัก (Application logic)
                    ↓
[ domain ] ← Entity + Interface
                    ↑
[ infrastructure (postgres) ] ← เชื่อมฐานข้อมูล

## 🚀 เริ่มต้นใช้งาน
### 1. สร้าง `.env`
### 2. `docker-compose up -d`
### 3. `go run ./cmd/main.go`

# 🧼 Go Clean Architecture - WARD-STOCK-BACKEND

โปรเจกต์นี้ใช้แนวทาง **Clean Architecture** ด้วยภาษา Go (Golang) โดยมีการแยกแต่ละเลเยอร์อย่างชัดเจน เพื่อให้โค้ดสามารถทดสอบง่าย ขยายต่อได้ และแยก concern อย่างเป็นระบบ

---

## 🔧 Stack ที่ใช้

- 🐹 Go 1.20+
- 🧱 GORM (ORM สำหรับเชื่อมต่อ PostgreSQL)
- 🌐 Gin (Web Framework)
- 🐘 PostgreSQL
- 🐳 Docker + docker-compose
- 📁 Clean Architecture Principles

---

## 📦 Folder Structure

WARD-STOCK-BACKEND/
├── cmd/ # Entry point (main.go)
├── internal/
│ ├── domain/ # Entity + Repository Interface
│ ├── usecase/ # Business logic (interfaces & impl)
│ ├── delivery/
│ │ └── http/ # HTTP handlers (Gin)
│ └── infrastructure/
│ └── postgres/ # DB access layer
├── model/ # (legacy or transitional structs)
├── repository/ # (optional: interface if not in domain)
├── .env, .env.development # Environment config
├── docker-compose.yml # For local container setup
└── go.mod / go.sum

## 🧱 Clean Architecture Layer Diagram

[ delivery (handler/http) ] ← Gin HTTP API
                    ↓
[ usecase ] ← ธุรกิจหลัก (Application logic)
                    ↓
[ domain ] ← Entity + Interface
                    ↑
[ infrastructure (postgres) ] ← เชื่อมฐานข้อมูล

## 🚀 เริ่มต้นใช้งาน
### 1. สร้าง `.env`
### 2. `docker-compose up -d`
### 3. `go run ./cmd/main.go`