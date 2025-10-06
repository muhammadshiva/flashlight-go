# Postman Collection Guide

## Files

1. **Flashlight_API.postman_collection.json** - Postman collection dengan semua endpoints
2. **Flashlight_API.postman_environment.json** - Environment variables untuk local development

## Cara Import ke Postman

### Import Collection

1. Buka Postman
2. Klik **Import** di pojok kiri atas
3. Pilih file `Flashlight_API.postman_collection.json`
4. Klik **Import**

### Import Environment (Optional tapi Recommended)

1. Klik **Import** di pojok kiri atas
2. Pilih file `Flashlight_API.postman_environment.json`
3. Klik **Import**
4. Pilih environment "Flashlight API - Local" dari dropdown di pojok kanan atas

## Cara Menggunakan Collection

### 1. Setup Environment Variables

Collection ini menggunakan environment variables berikut:

- `base_url` - URL server API (default: `http://localhost:8080`)
- `auth_token` - JWT token dari login (otomatis tersimpan setelah login)
- `user_id` - ID user (otomatis tersimpan setelah register)
- `work_order_id` - ID work order (otomatis tersimpan setelah create work order)

### 2. Testing Flow

Ikuti urutan ini untuk testing:

#### A. Health Check
```
GET /health
```
Pastikan API berjalan dengan baik.

#### B. Authentication Flow

1. **Register User**
   ```
   POST /api/v1/auth/register
   ```
   - Buat user baru
   - `user_id` akan otomatis tersimpan ke environment variable

2. **Login**
   ```
   POST /api/v1/auth/login
   ```
   - Login dengan credentials yang sama
   - `auth_token` akan otomatis tersimpan ke environment variable
   - Token ini akan digunakan untuk semua request selanjutnya

#### C. User Management

Semua endpoint users memerlukan authentication:

1. **Get All Users**
   ```
   GET /api/v1/users?page=1&per_page=10
   ```

2. **Get User by ID**
   ```
   GET /api/v1/users/{{user_id}}
   ```

3. **Update User**
   ```
   PUT /api/v1/users/{{user_id}}
   ```

4. **Delete User**
   ```
   DELETE /api/v1/users/{{user_id}}
   ```

#### D. Work Order Management

1. **Create Work Order**
   ```
   POST /api/v1/work-orders
   ```
   - `work_order_id` akan otomatis tersimpan

2. **Get All Work Orders**
   ```
   GET /api/v1/work-orders?page=1&per_page=10
   ```

3. **Get Work Orders by Status**
   ```
   GET /api/v1/work-orders?status=pending
   ```
   Status options: `pending`, `confirmed`, `in_progress`, `ready`, `completed`, `cancelled`

4. **Get Work Order by ID**
   ```
   GET /api/v1/work-orders/{{work_order_id}}
   ```

5. **Update Work Order**
   ```
   PUT /api/v1/work-orders/{{work_order_id}}
   ```

6. **Update Work Order Status** (Quick Actions)
   - Update to Confirmed
   - Update to In Progress
   - Update to Completed

7. **Delete Work Order**
   ```
   DELETE /api/v1/work-orders/{{work_order_id}}
   ```

#### E. Admin Operations

Requires role `owner` or `admin`:

1. **Create User (Admin Only)**
   ```
   POST /api/v1/admin/users
   ```

## Automatic Token Management

Collection ini sudah dilengkapi dengan **test scripts** yang otomatis menyimpan:

- JWT token setelah login ke `auth_token`
- User ID setelah register ke `user_id`
- Work Order ID setelah create ke `work_order_id`

Anda tidak perlu copy-paste token secara manual!

## Environment Variables Explanation

### base_url
Default: `http://localhost:8080`

Ubah jika server berjalan di port atau host yang berbeda:
```
http://localhost:3000
https://api.yourdomain.com
```

### auth_token
Terisi otomatis setelah login. Digunakan di header:
```
Authorization: Bearer {{auth_token}}
```

### user_id & work_order_id
Terisi otomatis setelah create, digunakan untuk testing endpoints by ID.

## Tips

1. **Pastikan server berjalan** sebelum testing
   ```bash
   make run
   ```

2. **Gunakan environment** untuk mudah switch antara local/staging/production

3. **Lihat console** di Postman untuk debug jika ada error

4. **Test scripts** akan otomatis save response data ke environment variables

5. **Disable query parameters** yang tidak dibutuhkan dengan checkbox di sebelah kiri

## Struktur Collection

```
Flashlight API
├── Health Check
│   └── Health Check
├── Authentication
│   ├── Register User
│   └── Login
├── Users
│   ├── Get All Users
│   ├── Get User by ID
│   ├── Update User
│   └── Delete User
├── Work Orders
│   ├── Create Work Order
│   ├── Get All Work Orders
│   ├── Get Work Orders by Status
│   ├── Get Work Order by ID
│   ├── Update Work Order
│   ├── Update Work Order Status to Confirmed
│   ├── Update Work Order Status to In Progress
│   ├── Update Work Order Status to Completed
│   └── Delete Work Order
└── Admin
    └── Create User (Admin Only)
```

## Troubleshooting

### Token Expired
Jika mendapat error `401 Unauthorized`:
1. Login ulang
2. Token akan otomatis di-refresh

### Invalid ID
Pastikan `user_id` atau `work_order_id` sudah terisi di environment variables.

### Connection Refused
Pastikan server API sudah berjalan:
```bash
make run
```

## Example Workflow

1. Start server: `make run`
2. Import collection & environment ke Postman
3. Pilih environment "Flashlight API - Local"
4. Run "Health Check" → pastikan status `ok`
5. Run "Register User" → `user_id` tersimpan
6. Run "Login" → `auth_token` tersimpan
7. Run "Create Work Order" → `work_order_id` tersimpan
8. Test semua endpoints lainnya!

Happy Testing! 🚀
