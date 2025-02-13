# RestApi-Golang-Employee

RestApi-Golang-Employee adalah proyek REST API berbasis Golang untuk mengelola data karyawan dengan MongoDB sebagai database utama.

## 📌 Teknologi yang Digunakan
- **Golang** 1.22.2
- **Gorilla Mux** (Router untuk HTTP request handling)
- **MongoDB Driver** (Integrasi dengan MongoDB)
- **UUID** (Pembuatan unique identifier)
- **Gotenv** (Manajemen variabel lingkungan)

## 📦 Instalasi

### 1. Clone Repository
```sh
git clone https://github.com/Dito-7/RestApi-Golang-Employee.git
cd RestApi-Golang-Employee
```

### 2. Install Dependencies
Pastikan Anda sudah menginstal Golang, lalu jalankan:
```sh
go mod tidy
```

### 3. Konfigurasi Environment
Buat file `.env` dan sesuaikan dengan konfigurasi MongoDB Anda:
```env
MONGO_URI=""
DB_NAME=""
COLLECTION_NAME=""
```

### 4. Jalankan Server
```sh
go run main.go
```
Server akan berjalan di `http://localhost:4444`

### 5. Api Dokumentasi
```sh
https://documenter.getpostman.com/view/33186673/2sAYXCkJic
```

## 📜 Lisensi
MIT License - bebas digunakan dan dikembangkan!

---
Jika ada pertanyaan atau ingin kontribusi, jangan ragu untuk membuat **issue** atau **pull request**! 😊
