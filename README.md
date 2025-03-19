# Inventory Management System

A modern web application built with Vue.js for managing warehouse inventory with real-time CRUD operations and authentication.

## Features

### Core Functionality

* 🛡️ JWT-based Authentication (Login/Register)
* 📦 Product Inventory Management (CRUD Operations)
* 🔍 Filter Products by Status and Stock Level
* 🏷️ Barcode Generation for Products
* 📥 CSV Export Functionality
* 📱 Responsive Mobile-first Design

## Project Setup

### Prerequisites

- Node.js (v18+)
- npm (v9+)
- Golang (v1.19+)
- MySQL

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/inventory-management-system.git
   ```
2. Install BE dependencies:

   ```bash
   go mod tidy
   ```
3. Configure environment variables (rename `.env.example` file to `.env` ) and change value under .env file
4. Run Backend server:

   ```bash
   go run /cmd/api/main.go
   ```
5. Change direktory to `/frontend` to run front end server:

   ```bash
   cd frontend
   ```
6. Install Front End dependencies:

   ```bash
   npm install
   ```
7. Configure environment variables (create `.env` file):

   ```env
   VITE_API_BASE_URL=http://localhost:8080
   ```
8. Run development server:

   ```bash
   npm run dev
   ```
9. d


## API Documentation

### Authentication

#### Register

```http
POST /api/v1/register
Content-Type: application/json

{
  "username": "admin",
  "password": "secret123",
  "email": "super@admin.com"
}

Response (200 OK):
{
    "id": "6bae33cd-2945-4a93-8385-3bb229456f65",
    "username": "admin",
    "email": "super@admin.com",
    "created_at": "2025-03-19T11:37:42.698017746+07:00",
    "updated_at": "2025-03-19T11:37:42.698017864+07:00"
}
```


#### Login

```http
POST /api/v1/login
Content-Type: application/json

{
  "username": "admin",
  "password": "secret123"
}

Response (200 OK):
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI0NDU1MjgsInVzZXJfaWQiOiI2YmFlMzNjZC0yOTQ1LTRhOTMtODM4NS0zYmIyMjk0NTZmNjUiLCJ1c2VybmFtZSI6Im9iYmllIn0.v-HUEZIpCn9n2OEdUDeIjoA7Ezr5WHIwHCJR249FsaM",
  "user": {
    	"id": "6bae33cd-2945-4a93-8385-3bb229456f65",
        "username": "admin",
        "email": "super@admin.com",
        "created_at": "2025-03-19T04:37:43Z",
        "updated_at": "2025-03-19T04:37:43Z"
  }
}
```

### Products

#### Get All Products

```http
GET /api/v1/products
Authorization: Bearer <token>

Response (200 OK):
[{
  "id": "5c44caeb-192c-434a-b388-d32eb7ef5577",
  "product_name": "Widget X",
  "sku": "WX-2023",
  "quantity": 42,
  "location": "Shelf A3",
  "status": "active"
}]
```

### Create Product

```http
POST /api/v1/products
Authorization: Bearer <token>
Content-Type: application/json

{
  "product_name": "Widget X",
  "sku": "WX-2023",
  "quantity": 42,
  "location": "Shelf A3",
  "status": "active"
}

Response (200 OK):
{
    "id": "5c44caeb-192c-434a-b388-d32eb7ef5577",
    "product_name": "Widget X",
    "sku": "WX-2023",
    "quantity": 42,
    "location": "Shelf A3",
    "status": "active"
    "created_at": "2025-03-19T11:48:22.619497784+07:00",
    "created_by": "6bae33cd-2945-4a93-8385-3bb229456f65",
    "updated_at": "2025-03-19T11:48:22.619497863+07:00",
    "updated_by": "6bae33cd-2945-4a93-8385-3bb229456f65"
}
```

### Detail Product

```http
POST /api/v1/products/5c44caeb-192c-434a-b388-d32eb7ef5577
Authorization: Bearer <token>

Response (200 OK):
{
    "id": "5c44caeb-192c-434a-b388-d32eb7ef5577",
    "product_name": "Widget X",
    "sku": "WX-2023",
    "quantity": 42,
    "location": "Shelf A3",
    "status": "active"
    "created_at": "2025-03-19T11:48:22.619497784+07:00",
    "created_by": "6bae33cd-2945-4a93-8385-3bb229456f65",
    "updated_at": "2025-03-19T11:48:22.619497863+07:00",
    "updated_by": "6bae33cd-2945-4a93-8385-3bb229456f65"
}
```

### Update Product

```http
PUT /api/v1/products/5c44caeb-192c-434a-b388-d32eb7ef5577
Authorization: Bearer <token>
Content-Type: application/json

{
  "product_name": "Widget X",
  "sku": "WX-2023",
  "quantity": 42,
  "location": "Shelf A3",
  "status": "active"
}

Response (200 OK):
{
    "id": "5c44caeb-192c-434a-b388-d32eb7ef5577",
    "product_name": "Widget X",
    "sku": "WX-2023",
    "quantity": 42,
    "location": "Shelf A3",
    "status": "active"
    "created_at": "2025-03-19T11:48:22.619497784+07:00",
    "created_by": "6bae33cd-2945-4a93-8385-3bb229456f65",
    "updated_at": "2025-03-19T11:48:22.619497863+07:00",
    "updated_by": "6bae33cd-2945-4a93-8385-3bb229456f65"
}
```

### Delete Product

```http
DELETE /api/v1/products/5c44caeb-192c-434a-b388-d32eb7ef5577
Authorization: Bearer <token>

Response (200 OK):
{
    "message": "Product deleted successfully"
}
```

### Export Products

```http
GET /api/v1/export/products
Authorization: Bearer <token>

Response:
200 OK (application/csv)
```


### Barcode Generation

```http
GET /api/v1/products/5c44caeb-192c-434a-b388-d32eb7ef5577/barcode
Authorization: Bearer <token>

Response:
200 OK (image/png)
```

## Screenshots

##### Register Screen
![Register Screen](screenshot/Register%20Screen.png)

##### Login Screen
![Login Screen](screenshot/Login%20Screen.png)

##### Create Product Screen
![Create Product Screen](screenshot/Create%20Product%20Screen.png)

##### Product List Screen
![List Products Screen](screenshot/List%20Products%20Screen.png)

##### Product Detail Screen
![Detail Product Screen](screenshot/Detail%20Product%20Screen.png)

##### Product Edit Screen
![Edit Product Screen](screenshot/Edit%20Product%20Screen.png)

##### Generate Barcode Screen
![Generate Barcode Screen](screenshot/Generate%20Barcode%20Screen.png)

##### Export To CSV Screen
![Export To CSV Screen](screenshot/Export%20To%20CSV%20Screen.png)

##### Delete Screen
![Delete Produk Screen](screenshot/Delete%20Produk%20Screen.png)
