---

# **Golang REST API Authentication with Fiber & GORM**  
Sebuah **REST API Authentication** menggunakan **Golang**, **Fiber**, dan **GORM** dengan fitur **Register, Login, Get Users**, dan **JWT Authentication**.  

## **🛠️ Tech Stack**
- **Golang** (Backend)
- **Fiber** (Web Framework)
- **GORM** (ORM untuk Database)
- **MySQL** (Database)
- **JWT** (JSON Web Token untuk Autentikasi)
- **Bcrypt** (Hashing Password)

---

## **📌 Fitur API**
✅ **Register User dengan Profile Default** 👤  
✅ **Login User dengan JWT Token** 🔐  
✅ **Get All Users** 📋  
✅ **Get User by ID (Beserta Profile)** 🔎  
✅ **JWT Authentication** 🔑  

---

## **⚙️ Setup & Menjalankan Secara Lokal**
### **1. Clone Repository**
```sh
git clone https://github.com/Adik-soleh/golang-rest-api-auth.git
cd golang-rest-api-auth
```

### **2. Install Dependencies**
```sh
go mod tidy
```

### **3. Konfigurasi Database**
- Pastikan **MySQL** berjalan di sistem kamu.
- Buat file **`.env`** di root project dan tambahkan:
  ```env
  DB_HOST=localhost
  DB_USER=root
  DB_PASS=password
  DB_NAME=golang_auth
  DB_PORT=3306
  JWT_SECRET=yoursecretkey
  ```

### **4. Jalankan Migration Database**
```sh
go run main.go
```
> **Catatan:** Jika tabel belum ada, GORM akan otomatis melakukan migrasi.

### **5. Jalankan Server**
```sh
go run main.go
```
> **Server berjalan di:** `http://localhost:3000`

---

## **📌 Endpoint API**
### **1️⃣ Register User**
- **Endpoint:** `POST /register`
- **Request Body:**
  ```json
  {
    "name": "Panjoel",
    "email": "panjoel@gmail.com",
    "password": "password123"
  }
  ```
- **Response Jika Sukses:**
  ```json
  {
    "message": "User registered successfully",
    "error": false,
    "data": {
      "id": 1,
      "name": "Panjoel",
      "email": "panjoel@gmail.com",
      "profile": {
        "id": 1,
        "username": "Panjoel",
        "bio": "No bio yet",
        "phone": "",
        "image": "default.png"
      }
    }
  }
  ```
- **Response Jika Email Sudah Terdaftar:**
  ```json
  {
    "message": "Email already exists",
    "error": true,
    "data": null
  }
  ```

### **2️⃣ Login User**
- **Endpoint:** `POST /login`
- **Request Body:**
  ```json
  {
    "email": "panjoel@gmail.com",
    "password": "password123"
  }
  ```
- **Response Jika Berhasil:**
  ```json
  {
    "message": "Login successful",
    "error": false,
    "data": {
      "token": "eyJhbGciOiJIUz..."
    }
  }
  ```
- **Response Jika Gagal (Email atau Password Salah):**
  ```json
  {
    "message": "Invalid credentials",
    "error": true,
    "data": null
  }
  ```

### **3️⃣ Get All Users**
- **Endpoint:** `GET /users`
- **Response:**
  ```json
  {
    "message": "Users retrieved successfully",
    "error": false,
    "data": [
      {
        "id": 1,
        "name": "Panjoel",
        "email": "panjoel@gmail.com",
        "profile": {
          "username": "Panjoel",
          "bio": "No bio yet",
          "phone": "",
          "image": "default.png"
        }
      }
    ]
  }
  ```

### **4️⃣ Get User by ID**
- **Endpoint:** `GET /users/:id`
- **Response Jika User Ada:**
  ```json
  {
    "message": "User retrieved successfully",
    "error": false,
    "data": {
      "id": 1,
      "name": "Panjoel",
      "email": "panjoel@gmail.com",
      "profile": {
        "username": "Panjoel",
        "bio": "No bio yet",
        "phone": "",
        "image": "default.png"
      }
    }
  }
  ```
- **Response Jika User Tidak Ada:**
  ```json
  {
    "message": "User not found",
    "error": true,
    "data": null
  }
  ```

---

## **🎯 To-Do List**
✅ **Register & Login dengan JWT**  
✅ **Hash Password dengan Bcrypt**  
✅ **Get All Users & Get User by ID (Beserta Profile)**  
⬜ Middleware JWT untuk proteksi endpoint 🔐  
⬜ Update & Delete User ✍️  

---

## **📜 Lisensi**
Proyek ini menggunakan lisensi **MIT** – bebas digunakan dan dikembangkan! 🚀  

> **Dibuat dengan ❤️ oleh [Adik Soleh]**  

---
