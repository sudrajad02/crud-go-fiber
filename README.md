# CRUD GO FIBER API

Crud Go Fiber project dengan fitur:
- CRUD Device
- CRUD Sensor

## **Persyaratan**
- Go Version go1.22.0
- Mysql

## **Instalasi**
1. Clone repo:
   ```bash
   git clone https://github.com/sudrajad02/crud-go-fiber
   cd crud-go-fiber
   ```
   
2. Tambahkan database jika belum ada di mysql
   ```bash
   CREATE DATABASE crud_go_fiber_db;
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Copy file `.env` dari `.env.sample` kemudian edit dengan sesuai dengan local yang sudah terinstall

5. Jalankan:
   ```bash
   go run main.go
   ```

## **Endpoint**
| Method | Endpoint         | Deskripsi             |
|--------|------------------|-----------------------|
| GET    | /api/device      | Ambil semua device    |
| POST   | /api/device      | Add new device        |
| PUT    | /api/device/:id  | Edit device           |
| DELETE | /api/device/:id  | Delete device         |
| GET    | /api/sensor      | Ambil semua sensor    |
| POST   | /api/sensor      | Add new sensor        |
| PUT    | /api/sensor/:id  | Edit sensor           |
| DELETE | /api/sensor/:id  | Delete sensor         |
