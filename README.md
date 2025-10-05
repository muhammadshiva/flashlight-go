# Flashlight GO - Work Order & POS Backend API

Backend API system menggunakan Golang untuk Platform Work Order dan Point of Sales (POS) Cashier. Sistem ini menyediakan REST API lengkap untuk manajemen pengguna, kendaraan, produk, work order, pembayaran, dan shift kasir.

## Fitur Utama

- ✅ **Manajemen User** - Multi-role (Owner, Admin, Cashier, Staff, Customer)
- ✅ **Membership System** - Tipe membership dengan benefits
- ✅ **Katalog Produk** - Kategori dan produk (Service, Addon, Retail)
- ✅ **Work Order** - Kelola order dari Kiosk, Cashier, atau Online
- ✅ **Pembayaran Multi-metode** - Cash, QRIS, Transfer, E-Wallet
- ✅ **Shift Management** - Tracking shift kasir dan total penjualan
- ✅ **Queue System** - Antrian otomatis untuk work order
- ✅ **FCM Push Notification** - Device token management

## Tech Stack

- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT (JSON Web Token)
- **Password Hashing**: bcrypt

## Struktur Project

```
flashlight-go/
├── cmd/
│   └── server/
│       └── main.go              # Entry point aplikasi
├── config/
│   └── config.go                # Konfigurasi aplikasi
├── internal/
│   ├── database/
│   │   └── database.go          # Database connection & migration
│   ├── dto/                     # Data Transfer Objects
│   │   ├── common_dto.go
│   │   ├── response.go
│   │   ├── user_dto.go
│   │   └── work_order_dto.go
│   ├── handler/                 # HTTP Handlers
│   │   ├── user_handler.go
│   │   └── work_order_handler.go
│   ├── middleware/              # HTTP Middlewares
│   │   ├── auth.go
│   │   └── cors.go
│   ├── models/                  # Database Models
│   │   ├── customer_vehicle.go
│   │   ├── device_fcm_token.go
│   │   ├── membership_type.go
│   │   ├── payment.go
│   │   ├── product.go
│   │   ├── product_category.go
│   │   ├── shift.go
│   │   ├── user.go
│   │   ├── vehicle.go
│   │   ├── work_order.go
│   │   └── work_order_item.go
│   ├── repository/              # Data Access Layer
│   │   ├── base_repository.go
│   │   ├── payment_repository.go
│   │   ├── repository.go
│   │   ├── shift_repository.go
│   │   ├── user_repository.go
│   │   └── work_order_repository.go
│   ├── routes/
│   │   └── routes.go            # Route definitions
│   └── service/                 # Business Logic Layer
│       ├── payment_service.go
│       ├── shift_service.go
│       ├── user_service.go
│       └── work_order_service.go
├── pkg/
│   └── utils/
│       └── jwt.go               # JWT utilities
├── .env.example                 # Environment variables template
├── .gitignore
├── go.mod
└── README.md
```

## Installation & Setup

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Git

### 1. Clone Repository

```bash
git clone <repository-url>
cd flashlight-go
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Setup Database

Buat database PostgreSQL:

```sql
CREATE DATABASE flashlight_db;
```

### 4. Configure Environment

Copy `.env.example` ke `.env` dan sesuaikan konfigurasi:

```bash
cp .env.example .env
```

Edit file `.env`:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=flashlight_db
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# JWT Configuration
JWT_SECRET=your-secret-key-change-this
JWT_EXPIRATION_HOURS=24

# Environment
APP_ENV=development
```

### 5. Run Application

```bash
go run cmd/server/main.go
```

Server akan berjalan di `http://localhost:8080`

## API Documentation

### Base URL

```
http://localhost:8080/api/v1
```

### Authentication

Semua endpoint yang protected memerlukan header:

```
Authorization: Bearer <jwt_token>
```

---

## API Endpoints

### Health Check

```http
GET /health
```

**Response:**

```json
{
  "status": "ok"
}
```

---

### Authentication

#### Register User

```http
POST /api/v1/auth/register
```

**Request Body:**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone_number": "08123456789",
  "password": "password123",
  "role": "customer",
  "address": "Jl. Contoh No. 123",
  "city": "Jakarta",
  "state": "DKI Jakarta",
  "postal_code": "12345",
  "country": "Indonesia"
}
```

**Response:**

```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "customer"
  }
}
```

#### Login

```http
POST /api/v1/auth/login
```

**Request Body:**

```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "customer"
    }
  }
}
```

---

### Users (Protected)

#### Get All Users

```http
GET /api/v1/users?page=1&per_page=10
```

#### Get User by ID

```http
GET /api/v1/users/:id
```

#### Update User

```http
PUT /api/v1/users/:id
```

**Request Body:**

```json
{
  "name": "John Doe Updated",
  "phone_number": "08199999999",
  "is_active": true
}
```

#### Delete User

```http
DELETE /api/v1/users/:id
```

---

### Work Orders (Protected)

#### Create Work Order

```http
POST /api/v1/work-orders
```

**Request Body:**

```json
{
  "source": "cashier",
  "type": "service",
  "customer_user_id": 1,
  "customer_vehicle_id": 1,
  "notes": "Cuci mobil premium",
  "items": [
    {
      "product_id": 1,
      "quantity": 1,
      "assigned_staff_user_id": 2,
      "item_note": "Extra wax"
    }
  ]
}
```

**Response:**

```json
{
  "success": true,
  "message": "Work order created successfully",
  "data": {
    "id": 1,
    "order_number": "WO-20250106-0001",
    "source": "cashier",
    "type": "service",
    "queue_number": 1,
    "status": "pending",
    "subtotal": 100000,
    "total_amount": 100000,
    "items": []
  }
}
```

#### Get All Work Orders

```http
GET /api/v1/work-orders?page=1&per_page=10
```

Filter by status:

```http
GET /api/v1/work-orders?status=pending
```

**Status values**: `pending`, `confirmed`, `in_progress`, `ready`, `completed`, `cancelled`

#### Get Work Order by ID

```http
GET /api/v1/work-orders/:id
```

#### Update Work Order

```http
PUT /api/v1/work-orders/:id
```

**Request Body:**

```json
{
  "status": "confirmed",
  "discount_amount": 10000,
  "tax_amount": 5000
}
```

#### Delete Work Order

```http
DELETE /api/v1/work-orders/:id
```

---

## Entitas Database

### 1. Users

- Multi-role: owner, admin, cashier, staff, customer
- Membership support
- Soft delete

### 2. Membership Types

- Flexible benefits (JSON)
- Active/inactive status

### 3. Vehicles & Customer Vehicles

- Master data kendaraan
- Relasi customer dengan kendaraan
- License plate tracking

### 4. Product Categories & Products

- Kategori produk dengan icon
- Product kind: service, addon, retail
- Premium product flag

### 5. Work Orders

- Multi-source: kiosk, cashier, online
- Status workflow lengkap
- Queue number otomatis
- Financial tracking (subtotal, discount, tax, total)

### 6. Work Order Items

- Product snapshot (nama & harga)
- Staff assignment
- Item notes

### 7. Payments

- Multi-method: cash, qris, transfer, e_wallet
- Multi-payment support (DP, cicilan)
- Change calculation
- Payment status tracking

### 8. Shifts

- Kasir shift management
- Initial & final cash tracking
- Sales summary

### 9. Device FCM Tokens

- Push notification support
- Multi-device per user

---

## Alur Bisnis Utama

### 1. Kiosk Flow

1. Customer membuat work order di kiosk (source=kiosk)
2. Pilih kendaraan & paket service
3. Sistem assign queue number
4. Status: pending → confirmed → in_progress → ready → completed

### 2. Cashier Flow

1. Kasir buka shift (start shift)
2. Buat/kelola work orders
3. Proses pembayaran (tunai/non-tunai)
4. Tutup shift dengan final cash count

### 3. Payment Flow

1. Work order bisa dibayar bertahap (DP, pelunasan)
2. Sistem track total pembayaran
3. Auto-complete work order saat fully paid

### 4. Retail Flow

1. Buat work order dengan type=retail
2. Tambahkan produk retail
3. Langsung bayar dan selesai

---

## Role-Based Access Control

- **Owner**: Full access
- **Admin**: Manajemen sistem
- **Cashier**: Work orders, payments, shifts
- **Staff**: Assigned work order items
- **Customer**: View own orders

---

## Development

### Run in Development Mode

```bash
go run cmd/server/main.go
```

### Build for Production

```bash
go build -o flashlight-go cmd/server/main.go
./flashlight-go
```

---

## Environment Variables

| Variable               | Description       | Default       |
| ---------------------- | ----------------- | ------------- |
| `DB_HOST`              | Database host     | localhost     |
| `DB_PORT`              | Database port     | 5432          |
| `DB_USER`              | Database user     | postgres      |
| `DB_PASSWORD`          | Database password | postgres      |
| `DB_NAME`              | Database name     | flashlight_db |
| `DB_SSLMODE`           | SSL mode          | disable       |
| `SERVER_PORT`          | Server port       | 8080          |
| `SERVER_HOST`          | Server host       | 0.0.0.0       |
| `JWT_SECRET`           | JWT secret key    | -             |
| `JWT_EXPIRATION_HOURS` | JWT expiration    | 24            |
| `APP_ENV`              | Environment       | development   |

---

## Contributing

1. Fork repository
2. Create feature branch
3. Commit changes
4. Push to branch
5. Open Pull Request

---

## License

Copyright © 2025 Matariza
