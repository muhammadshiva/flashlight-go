---

## Mermaid ERD

```mermaid
erDiagram
    %% Users & Membership
    users ||--o{ device_fcm_tokens : "has many"
    users }o--|| membership_types : "optional membership"
    users ||--o{ customer_vehicles : "owns many"
    users ||--o{ work_orders : "creates/handled by"
    users ||--o{ payments : "processed by (cashier)"

    %% Catalog
    product_categories ||--o{ products : "has many"

    %% Orders
    customer_vehicles ||--o{ work_orders : "used in"
    work_orders ||--o{ work_order_items : "contains"
    work_orders ||--o{ payments : "paid by"

    %% Shifts
    shifts ||--o{ work_orders : "logged in shift"
    shifts ||--o{ payments : "logged in shift"

    %% USERS
    users {
        bigint id PK
        string name
        string email UK
        string phone_number
        string password
        enum role "owner,admin,cashier,staff,customer"
        bigint membership_type_id FK NULL
        datetime membership_expires_at NULL
        string address NULL
        string city NULL
        string state NULL
        string postal_code NULL
        string country NULL
        string profile_image NULL
        boolean is_active DEFAULT true
        datetime last_login_at NULL
        string fcm_token NULL
        datetime email_verified_at NULL
        string remember_token NULL
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at NULL
    }

    %% MEMBERSHIP
    membership_types {
        bigint id PK
        string name
        json benefits NULL
        boolean is_active DEFAULT true
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at NULL
    }

    %% DEVICES (Push)
    device_fcm_tokens {
        bigint id PK
        bigint user_id FK
        string device_id UK
        string fcm_token
        string device_name NULL
        string device_type NULL
        timestamp last_used_at NULL
        timestamp created_at
        timestamp updated_at
    }

    %% VEHICLES
    vehicles {
        bigint id PK
        string brand
        string model
        string vehicle_type
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    %% CUSTOMER VEHICLES
    customer_vehicles {
        bigint id PK
        bigint customer_id FK
        bigint vehicle_id FK
        string license_plate
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    %% CATALOG
    product_categories {
        bigint id PK
        string name
        string icon_image NULL
        boolean is_active DEFAULT true
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at NULL
    }

    %% PRODUCTS
    products {
        bigint id PK
        string name
        text description NULL
        decimal price
        string image NULL
        bigint category_id FK
        enum kind "service,addon,retail"
        boolean is_active DEFAULT true
        boolean is_premium DEFAULT false
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at NULL
    }

    %% ORDER 
    work_orders {
        bigint id PK
        string order_number UK
        enum source "kiosk,cashier,online"
        enum type "service,retail,mix"
        bigint customer_user_id FK NULL
        bigint customer_vehicle_id FK NULL
        bigint cashier_user_id FK NULL
        bigint shift_id FK NULL
        integer queue_number NULL
        enum status "pending,confirmed,in_progress,ready,completed,cancelled"
        text notes NULL
        text special_instructions NULL
        datetime confirmed_at NULL
        datetime started_at NULL
        datetime completed_at NULL

        decimal subtotal DEFAULT 0
        decimal discount_amount DEFAULT 0
        decimal tax_amount DEFAULT 0
        decimal total_amount DEFAULT 0

        timestamp created_at
        timestamp updated_at
        timestamp deleted_at NULL
    }

    work_order_items {
        bigint id PK
        bigint work_order_id FK
        bigint product_id FK
        string product_name_snapshot
        decimal price_snapshot
        integer quantity
        decimal subtotal

        bigint assigned_staff_user_id FK NULL
        text item_note NULL

        timestamp created_at
        timestamp updated_at
    }

    %% Pembayaran multi-termin (DP, split method)
    payments {
        bigint id PK
        bigint work_order_id FK
        bigint cashier_user_id FK NULL
        bigint shift_id FK NULL

        string payment_number UK
        enum method "cash,qris,transfer,e_wallet"
        enum status "pending,completed,failed,refunded"
        decimal amount_paid
        decimal change_amount DEFAULT 0
        string reference_number NULL
        json raw_payload NULL
        datetime paid_at NULL

        timestamp created_at
        timestamp updated_at
    }

    %% Shift Kasir
    shifts {
        bigint id PK
        bigint user_id FK
        datetime start_time
        datetime end_time NULL
        decimal initial_cash DEFAULT 0
        decimal final_cash DEFAULT 0
        decimal total_sales DEFAULT 0
        enum status "active,closed,canceled"
        string received_from NULL
        timestamp created_at
        timestamp updated_at
    }
```

## Indeks yang Disarankan
- `work_orders(order_number)`, `work_orders(customer_user_id)`, `work_orders(shift_id)`, `work_orders(status, source, created_at)`
- `work_order_items(work_order_id)`, `work_order_items(product_id)`
- `payments(work_order_id)`, `payments(payment_number)`, `payments(shift_id)`
- `customer_vehicles(customer_user_id, license_plate)`
- `products(category_id, kind, is_active)`

---

## Alur Singkat yang Didukung
1. **Kiosk/Front Desk** membuat **work_order (source=kiosk)** → pilih kendaraan & paket → sistem memberi **queue_number**.  
2. **Teknisi/Washer** memproses: `pending/confirmed → in_progress → ready → completed`.  
3. **Kasir** menutup transaksi: input **payments** (DP/pelunasan/split) → status `completed` saat bayar ≥ total.  
4. **Retail-only**: buat **work_order (type=retail)** + item retail → langsung **payments**.  
5. **Shift**: semua **work_order** & **payments** tercatat ke `shift_id` kasir aktif.
