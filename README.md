# 🧼  WARD-STOCK-BACKEND

---

## 🔧 Stack ที่ใช้

- 🐹 Go 1.20+
- 🧱 GORM (ORM สำหรับเชื่อมต่อ PostgreSQL)
- 🌐 Gin (Web Framework)
- 🐘 PostgreSQL
- 🐳 Docker + docker-compose
- 📁 Clean Architecture Principles

---

## 🚀 เริ่มต้นใช้งาน
### 1. สร้าง `.env`

สร้างไฟล์ `.env.development` (หรือ `.env`) และใส่ค่าดังนี้:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=ward_stock_dev_user
DB_PASSWORD=MvJCAy0dp24lcnkRMrtSKOovcvXs0J0J
DB_NAME=ward_stock_dev
DB_SSLMODE=disable

JWT_SECRET=your-super-secret-key


### 2. `docker-compose up -d`
### 3. `make run` // สร้าง Table
### 4. `make seed` // สร้าง Seed Test ข้อมูล

---


26 internal/delivery/http/user_handler.go เปิดการ Auth


---

make run        # รัน server
make seed       # เติมข้อมูล test
make swag       # สร้าง Swagger doc ใหม่
make clean-docs # ล้าง Swagger doc เก่า
