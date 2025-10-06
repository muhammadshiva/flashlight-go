# Flashlight API Documentation

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication
Most endpoints require authentication using JWT token. Include the token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

---

## How to Run the API

### Prerequisites
1. Go 1.x or higher installed
2. PostgreSQL database installed and running
3. Environment variables configured (see `.env.example`)

### Setup Steps

1. **Clone and navigate to the project**
```bash
cd /path/to/flashlight-go
```

2. **Install dependencies**
```bash
make deps
# or
go mod download
go mod tidy
```

3. **Setup environment variables**
```bash
cp .env.example .env
# Edit .env with your database credentials
```

4. **Create database**
```bash
make db-setup
# or manually
createdb flashlight_db
```

5. **Run the server**
```bash
make run
# or
go run cmd/server/main.go
```

The server will start on `http://localhost:8080` (or the port specified in your `.env` file)

### Additional Commands

- **Build binary**: `make build` or `go build -o flashlight-go cmd/server/main.go`
- **Run built binary**: `./flashlight-go`
- **Format code**: `make fmt`
- **Clean build artifacts**: `make clean`

---

## Endpoints

### Health Check

#### GET /health
Check if the API is running.

**Authentication**: Not required

**Response**:
```json
{
  "status": "ok"
}
```

---

## Authentication Endpoints

### 1. Register User

#### POST /api/v1/auth/register
Create a new user account.

**Authentication**: Not required

**Request Body**:
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone_number": "+1234567890",
  "password": "securePassword123",
  "role": "customer",
  "membership_type_id": 1,
  "membership_expires_at": "2024-12-31T23:59:59Z",
  "address": "123 Main St",
  "city": "New York",
  "state": "NY",
  "postal_code": "10001",
  "country": "USA"
}
```

**Required Fields**:
- `name` (string)
- `email` (string, valid email format)
- `password` (string, minimum 6 characters)
- `role` (string, one of: `owner`, `admin`, `cashier`, `staff`, `customer`)

**Optional Fields**:
- `phone_number` (string)
- `membership_type_id` (uint)
- `membership_expires_at` (string, ISO 8601 format)
- `address` (string)
- `city` (string)
- `state` (string)
- `postal_code` (string)
- `country` (string)

**Success Response** (201 Created):
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone_number": "+1234567890",
    "role": "customer",
    "membership_type_id": 1,
    "membership_expires_at": "2024-12-31T23:59:59Z",
    "address": "123 Main St",
    "city": "New York",
    "state": "NY",
    "postal_code": "10001",
    "country": "USA",
    "profile_image": null,
    "is_active": true,
    "last_login_at": null,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

**Error Response** (400 Bad Request):
```json
{
  "success": false,
  "message": "Invalid request",
  "error": "validation error details"
}
```

---

### 2. Login

#### POST /api/v1/auth/login
Authenticate user and receive JWT token.

**Authentication**: Not required

**Request Body**:
```json
{
  "email": "john.doe@example.com",
  "password": "securePassword123"
}
```

**Required Fields**:
- `email` (string, valid email format)
- `password` (string)

**Success Response** (200 OK):
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "phone_number": "+1234567890",
      "role": "customer",
      "membership_type_id": 1,
      "membership_expires_at": "2024-12-31T23:59:59Z",
      "address": "123 Main St",
      "city": "New York",
      "state": "NY",
      "postal_code": "10001",
      "country": "USA",
      "profile_image": null,
      "is_active": true,
      "last_login_at": "2024-01-01T10:00:00Z",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  }
}
```

**Error Response** (401 Unauthorized):
```json
{
  "success": false,
  "message": "Login failed",
  "error": "invalid credentials"
}
```

---

## User Endpoints

All user endpoints (except register) require authentication.

### 3. Get All Users

#### GET /api/v1/users
Retrieve a paginated list of users.

**Authentication**: Required

**Query Parameters**:
- `page` (optional, default: 1)
- `per_page` (optional, default: 10)

**Example Request**:
```
GET /api/v1/users?page=1&per_page=10
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "phone_number": "+1234567890",
      "role": "customer",
      "membership_type_id": 1,
      "membership_expires_at": "2024-12-31T23:59:59Z",
      "address": "123 Main St",
      "city": "New York",
      "state": "NY",
      "postal_code": "10001",
      "country": "USA",
      "profile_image": null,
      "is_active": true,
      "last_login_at": "2024-01-01T10:00:00Z",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "meta": {
    "page": 1,
    "per_page": 10,
    "total": 50,
    "total_pages": 5
  }
}
```

---

### 4. Get User by ID

#### GET /api/v1/users/:id
Retrieve a specific user by ID.

**Authentication**: Required

**URL Parameters**:
- `id` (required, uint)

**Example Request**:
```
GET /api/v1/users/1
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone_number": "+1234567890",
    "role": "customer",
    "membership_type_id": 1,
    "membership_expires_at": "2024-12-31T23:59:59Z",
    "address": "123 Main St",
    "city": "New York",
    "state": "NY",
    "postal_code": "10001",
    "country": "USA",
    "profile_image": null,
    "is_active": true,
    "last_login_at": "2024-01-01T10:00:00Z",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

**Error Response** (404 Not Found):
```json
{
  "success": false,
  "message": "User not found",
  "error": "record not found"
}
```

---

### 5. Update User

#### PUT /api/v1/users/:id
Update an existing user.

**Authentication**: Required

**URL Parameters**:
- `id` (required, uint)

**Request Body** (all fields optional):
```json
{
  "name": "Jane Doe",
  "email": "jane.doe@example.com",
  "phone_number": "+0987654321",
  "password": "newPassword123",
  "role": "admin",
  "membership_type_id": 2,
  "membership_expires_at": "2025-12-31T23:59:59Z",
  "address": "456 Oak Ave",
  "city": "Los Angeles",
  "state": "CA",
  "postal_code": "90001",
  "country": "USA",
  "is_active": true
}
```

**All Fields Optional**:
- `name` (string)
- `email` (string, valid email format)
- `phone_number` (string)
- `password` (string, minimum 6 characters)
- `role` (string, one of: `owner`, `admin`, `cashier`, `staff`, `customer`)
- `membership_type_id` (uint)
- `membership_expires_at` (string, ISO 8601 format)
- `address` (string)
- `city` (string)
- `state` (string)
- `postal_code` (string)
- `country` (string)
- `is_active` (boolean)

**Success Response** (200 OK):
```json
{
  "success": true,
  "message": "User updated successfully",
  "data": {
    "id": 1,
    "name": "Jane Doe",
    "email": "jane.doe@example.com",
    "phone_number": "+0987654321",
    "role": "admin",
    "membership_type_id": 2,
    "membership_expires_at": "2025-12-31T23:59:59Z",
    "address": "456 Oak Ave",
    "city": "Los Angeles",
    "state": "CA",
    "postal_code": "90001",
    "country": "USA",
    "profile_image": null,
    "is_active": true,
    "last_login_at": "2024-01-01T10:00:00Z",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-02T15:30:00Z"
  }
}
```

---

### 6. Delete User

#### DELETE /api/v1/users/:id
Delete a user (soft delete).

**Authentication**: Required

**URL Parameters**:
- `id` (required, uint)

**Example Request**:
```
DELETE /api/v1/users/1
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "message": "User deleted successfully",
  "data": null
}
```

**Error Response** (500 Internal Server Error):
```json
{
  "success": false,
  "message": "Failed to delete user",
  "error": "error details"
}
```

---

## Work Order Endpoints

All work order endpoints require authentication.

### 7. Create Work Order

#### POST /api/v1/work-orders
Create a new work order.

**Authentication**: Required

**Request Body**:
```json
{
  "source": "cashier",
  "type": "service",
  "customer_user_id": 5,
  "customer_vehicle_id": 3,
  "notes": "Customer requested oil change and tire rotation",
  "special_instructions": "Use synthetic oil",
  "items": [
    {
      "product_id": 10,
      "quantity": 1,
      "assigned_staff_user_id": 7,
      "item_note": "Check for leaks"
    },
    {
      "product_id": 15,
      "quantity": 4,
      "assigned_staff_user_id": 8,
      "item_note": "Rotate tires in X pattern"
    }
  ]
}
```

**Required Fields**:
- `source` (string, one of: `kiosk`, `cashier`, `online`)
- `type` (string, one of: `service`, `retail`, `mix`)
- `items` (array, minimum 1 item)
  - `product_id` (uint, required)
  - `quantity` (int, required, minimum 1)

**Optional Fields**:
- `customer_user_id` (uint)
- `customer_vehicle_id` (uint)
- `notes` (string)
- `special_instructions` (string)
- Per item:
  - `assigned_staff_user_id` (uint)
  - `item_note` (string)

**Success Response** (201 Created):
```json
{
  "success": true,
  "message": "Work order created successfully",
  "data": {
    "id": 1,
    "order_number": "WO-20240101-0001",
    "source": "cashier",
    "type": "service",
    "customer_user_id": 5,
    "customer_vehicle_id": 3,
    "cashier_user_id": 2,
    "shift_id": null,
    "queue_number": 1,
    "status": "pending",
    "notes": "Customer requested oil change and tire rotation",
    "special_instructions": "Use synthetic oil",
    "confirmed_at": null,
    "started_at": null,
    "completed_at": null,
    "subtotal": 89.98,
    "discount_amount": 0,
    "tax_amount": 7.20,
    "total_amount": 97.18,
    "items": [
      {
        "id": 1,
        "work_order_id": 1,
        "product_id": 10,
        "product_name_snapshot": "Oil Change - Synthetic",
        "price_snapshot": 49.99,
        "quantity": 1,
        "subtotal": 49.99,
        "assigned_staff_user_id": 7,
        "item_note": "Check for leaks",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      },
      {
        "id": 2,
        "work_order_id": 1,
        "product_id": 15,
        "product_name_snapshot": "Tire Rotation",
        "price_snapshot": 39.99,
        "quantity": 1,
        "subtotal": 39.99,
        "assigned_staff_user_id": 8,
        "item_note": "Rotate tires in X pattern",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      }
    ],
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

---

### 8. Get All Work Orders

#### GET /api/v1/work-orders
Retrieve a paginated list of work orders.

**Authentication**: Required

**Query Parameters**:
- `page` (optional, default: 1)
- `per_page` (optional, default: 10)
- `status` (optional, filters by status: `pending`, `confirmed`, `in_progress`, `ready`, `completed`, `cancelled`)

**Example Requests**:
```
GET /api/v1/work-orders?page=1&per_page=10
GET /api/v1/work-orders?status=pending
GET /api/v1/work-orders?status=in_progress&page=2
```

**Success Response with Pagination** (200 OK):
```json
{
  "success": true,
  "message": "Work orders retrieved successfully",
  "data": [
    {
      "id": 1,
      "order_number": "WO-20240101-0001",
      "source": "cashier",
      "type": "service",
      "customer_user_id": 5,
      "customer_vehicle_id": 3,
      "cashier_user_id": 2,
      "shift_id": null,
      "queue_number": 1,
      "status": "pending",
      "notes": "Customer requested oil change",
      "special_instructions": "Use synthetic oil",
      "confirmed_at": null,
      "started_at": null,
      "completed_at": null,
      "subtotal": 49.99,
      "discount_amount": 0,
      "tax_amount": 4.00,
      "total_amount": 53.99,
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "meta": {
    "page": 1,
    "per_page": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

**Success Response with Status Filter** (200 OK):
```json
{
  "success": true,
  "message": "Work orders retrieved successfully",
  "data": [
    {
      "id": 2,
      "order_number": "WO-20240101-0002",
      "source": "online",
      "type": "retail",
      "customer_user_id": 8,
      "customer_vehicle_id": null,
      "cashier_user_id": null,
      "shift_id": null,
      "queue_number": 2,
      "status": "pending",
      "notes": null,
      "special_instructions": null,
      "confirmed_at": null,
      "started_at": null,
      "completed_at": null,
      "subtotal": 29.99,
      "discount_amount": 5.00,
      "tax_amount": 2.00,
      "total_amount": 26.99,
      "created_at": "2024-01-01T11:00:00Z",
      "updated_at": "2024-01-01T11:00:00Z"
    }
  ]
}
```

---

### 9. Get Work Order by ID

#### GET /api/v1/work-orders/:id
Retrieve a specific work order by ID with all items.

**Authentication**: Required

**URL Parameters**:
- `id` (required, uint)

**Example Request**:
```
GET /api/v1/work-orders/1
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "message": "Work order retrieved successfully",
  "data": {
    "id": 1,
    "order_number": "WO-20240101-0001",
    "source": "cashier",
    "type": "service",
    "customer_user_id": 5,
    "customer_vehicle_id": 3,
    "cashier_user_id": 2,
    "shift_id": null,
    "queue_number": 1,
    "status": "in_progress",
    "notes": "Customer requested oil change and tire rotation",
    "special_instructions": "Use synthetic oil",
    "confirmed_at": "2024-01-01T10:05:00Z",
    "started_at": "2024-01-01T10:15:00Z",
    "completed_at": null,
    "subtotal": 89.98,
    "discount_amount": 0,
    "tax_amount": 7.20,
    "total_amount": 97.18,
    "items": [
      {
        "id": 1,
        "work_order_id": 1,
        "product_id": 10,
        "product_name_snapshot": "Oil Change - Synthetic",
        "price_snapshot": 49.99,
        "quantity": 1,
        "subtotal": 49.99,
        "assigned_staff_user_id": 7,
        "item_note": "Check for leaks",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      },
      {
        "id": 2,
        "work_order_id": 1,
        "product_id": 15,
        "product_name_snapshot": "Tire Rotation",
        "price_snapshot": 39.99,
        "quantity": 1,
        "subtotal": 39.99,
        "assigned_staff_user_id": 8,
        "item_note": "Rotate tires in X pattern",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      }
    ],
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:15:00Z"
  }
}
```

**Error Response** (404 Not Found):
```json
{
  "success": false,
  "message": "Work order not found",
  "error": "record not found"
}
```

---

### 10. Update Work Order

#### PUT /api/v1/work-orders/:id
Update an existing work order.

**Authentication**: Required

**URL Parameters**:
- `id` (required, uint)

**Request Body** (all fields optional):
```json
{
  "status": "completed",
  "cashier_user_id": 2,
  "shift_id": 5,
  "notes": "Updated notes",
  "special_instructions": "Updated instructions",
  "discount_amount": 10.00,
  "tax_amount": 6.50
}
```

**All Fields Optional**:
- `status` (string, one of: `pending`, `confirmed`, `in_progress`, `ready`, `completed`, `cancelled`)
- `cashier_user_id` (uint)
- `shift_id` (uint)
- `notes` (string)
- `special_instructions` (string)
- `discount_amount` (float64)
- `tax_amount` (float64)

**Success Response** (200 OK):
```json
{
  "success": true,
  "message": "Work order updated successfully",
  "data": {
    "id": 1,
    "order_number": "WO-20240101-0001",
    "source": "cashier",
    "type": "service",
    "customer_user_id": 5,
    "customer_vehicle_id": 3,
    "cashier_user_id": 2,
    "shift_id": 5,
    "queue_number": 1,
    "status": "completed",
    "notes": "Updated notes",
    "special_instructions": "Updated instructions",
    "confirmed_at": "2024-01-01T10:05:00Z",
    "started_at": "2024-01-01T10:15:00Z",
    "completed_at": "2024-01-01T11:30:00Z",
    "subtotal": 89.98,
    "discount_amount": 10.00,
    "tax_amount": 6.50,
    "total_amount": 86.48,
    "items": [
      {
        "id": 1,
        "work_order_id": 1,
        "product_id": 10,
        "product_name_snapshot": "Oil Change - Synthetic",
        "price_snapshot": 49.99,
        "quantity": 1,
        "subtotal": 49.99,
        "assigned_staff_user_id": 7,
        "item_note": "Check for leaks",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      },
      {
        "id": 2,
        "work_order_id": 1,
        "product_id": 15,
        "product_name_snapshot": "Tire Rotation",
        "price_snapshot": 39.99,
        "quantity": 1,
        "subtotal": 39.99,
        "assigned_staff_user_id": 8,
        "item_note": "Rotate tires in X pattern",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      }
    ],
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T11:30:00Z"
  }
}
```

---

### 11. Delete Work Order

#### DELETE /api/v1/work-orders/:id
Delete a work order (soft delete).

**Authentication**: Required

**URL Parameters**:
- `id` (required, uint)

**Example Request**:
```
DELETE /api/v1/work-orders/1
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "message": "Work order deleted successfully",
  "data": null
}
```

**Error Response** (500 Internal Server Error):
```json
{
  "success": false,
  "message": "Failed to delete work order",
  "error": "error details"
}
```

---

## Admin Endpoints

Admin endpoints require authentication and specific roles (owner or admin).

### 12. Create User (Admin)

#### POST /api/v1/admin/users
Create a new user (admin only).

**Authentication**: Required (Role: owner or admin)

**Request Body**: Same as [Register User](#1-register-user)

**Success Response**: Same as [Register User](#1-register-user)

---

## Status Codes

- `200 OK`: Request succeeded
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request format or validation error
- `401 Unauthorized`: Authentication failed or token missing
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

---

## Response Format

All API responses follow this standard format:

### Success Response
```json
{
  "success": true,
  "message": "Operation description",
  "data": { ... }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Error description",
  "error": "Detailed error message"
}
```

### Paginated Response
```json
{
  "success": true,
  "message": "Operation description",
  "data": [ ... ],
  "meta": {
    "page": 1,
    "per_page": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

---

## Example Usage with cURL

### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "securePassword123"
  }'
```

### Get All Users (with authentication)
```bash
curl -X GET http://localhost:8080/api/v1/users?page=1&per_page=10 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Create Work Order
```bash
curl -X POST http://localhost:8080/api/v1/work-orders \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "source": "cashier",
    "type": "service",
    "customer_user_id": 5,
    "items": [
      {
        "product_id": 10,
        "quantity": 1
      }
    ]
  }'
```

### Update Work Order Status
```bash
curl -X PUT http://localhost:8080/api/v1/work-orders/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "completed"
  }'
```

---

## Notes

- All timestamps are in ISO 8601 format (UTC)
- The API uses JWT tokens for authentication with a default expiration of 24 hours (configurable via environment variables)
- Passwords must be at least 6 characters long
- User roles: `owner`, `admin`, `cashier`, `staff`, `customer`
- Work order sources: `kiosk`, `cashier`, `online`
- Work order types: `service`, `retail`, `mix`
- Work order statuses: `pending`, `confirmed`, `in_progress`, `ready`, `completed`, `cancelled`
