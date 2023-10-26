### simple-digital-wallet ( TEST MNC 2023 )

## Deskripsi

Ini adalah sebuah test yang diberikan oleh MNC

## Dokumentasi Spesifikasi API

Selamat datang dalam dokumentasi spesifikasi API proyek ini. Dokumen ini memberikan panduan tentang cara menggunakan dan mengakses API yang disediakan dalam proyek ini.

## Daftar Isi

- [Deskripsi](#deskripsi)
- [Dokumentasi Spesifikasi API](#dokumentasi-spesifikasi-api)
- [Daftar Isi](#daftar-isi)
- [Deskripsi API](#deskripsi-api)
- [Endpoints API](#endpoints-api)
- [Penggunaan API](#penggunaan-api)
- [Instalasi](#instalasi)
- [Otentikasi](#otentikasi)
- [Contoh Permintaan (Request)](#contoh-permintaan-request)
  - [Auth API](#auth-api)
  - [POST Registration](#post-registration)
  - [POST Login](#post-login)
  - [POST Logout](#post-logout)
  - [Transaction API](#transaction-api)
  - [POST Transaction](#post-transaction)
  - [GET Transaction History](#get-transaction-history)
  - [Customer API](#customer-api)
  - [GET Profile](#get-profile)

## Deskripsi API

API ini menyediakan akses ke berbagai layanan dan fitur dalam proyek ini. Anda dapat menggunakan API ini untuk berinteraksi dengan data dan sumber daya yang disediakan oleh proyek.

## Endpoints API

Berikut adalah beberapa contoh endpoint yang tersedia dalam API ini:

- `/customers`: Mengelola user dalam produk ini.
- `/auth`: Mengelola autentikasi dan autorisasi user.
- `/transactions`: Mengelola Transaksi.

## Penggunaan API

Untuk menggunakan API ini, Anda perlu melakukan permintaan HTTP ke endpoint yang sesuai. Pastikan Anda telah menyiapkan otorisasi atau otentikasi jika diperlukan.

## Instalasi

1. Clone repositori ini:

   ```bash
   git clone https://github.com/cypherpunkx/test-mnc-2023.git
   cd test-mnc-2023
   go mod tidy
   go run main.go
   ```

## Otentikasi

Beberapa permintaan ke API ini mungkin memerlukan otentikasi. Anda dapat menggunakan token API atau metode otentikasi lainnya sesuai dengan petunjuk yang ada. Token API biasanya disediakan saat Anda mendaftar atau masuk sebagai pengguna.

## Contoh Permintaan (Request)

### Auth API

### POST Registration
Request :

- Method : POST
- Endpoint : `/auth/registration`
- Header :
  - Accept : application/json
- Body :

```json
{
  "userName": "john_doe",
  "email": "johndoe@example.com",
  "password": "securepassword",
  "firstName": "John",
  "lastName": "Doe",
  "bankCard": {
  "bankName": "bca",
  "cardNumber": "123456",
  "cardholderName": "John Doe",
  "expirationDate": "2024-09-30T00:00:00Z",
  "balance": 1000.0,
  "cvv": "123"
  }
}
```

Response :

- Status : 201 CREATED
- Body :

```json
{
    "code":201,
    "status":"success",
    "message":"Registration Successful!"
}

```
### POST Login
Request :

- Method : POST
- Endpoint : `/auth/login`
- Header :
  - Accept : application/json
- Body :
  
```json
{
  "userName": "john_doe",
  "password": "securepassword",
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "code":200,
    "status":"success",
    "message":"Login Successful!",
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

### POST Logout
Request :

- Method : POST
- Endpoint : `/auth/logout`
- Header :
  - Accept : application/json
  - Authorization: Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
    "code":200,
    "status":"success",
    "message":"Logged out Successful!"
}e
```

### Transaction API

### POST Transaction
Request :

- Method : POST
- Endpoint : `/customers/transactions/:id/send`
- Header :
  - Accept : application/json
  - Authorization : Bearer Token
- Body :

```json
{
  "transactionType": "send",
  "amount": 100.5,
  "description": "sending money to other customer or friend",
  "timestamp": "2023-09-20T15:30:00Z"
}

```

Response :

- Status : 201 CREATED
- Body :

```json
{
 "code":201,
 "status":"success",
 "message":"Transaction Successful!",
 "data": {
     "id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
  "UserID": "4d7c52eb-7d11-4826-9c97-6bea4fb2ae62",
  "transactionType": "send",
  "amount": 100.5,
 "description": "sending money to other customer or friend",
  "timestamp": "2023-09-20T15:30:00Z"
 }
}

```

### GET Transaction History
Request :

- Method : GET
- Endpoint : `/customers/transactions/history`
- Header :
  - Accept : application/json
  - Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
 "code":200,
 "status":"Success",
 "message":"Get All Transactions",
  "data": {
  "id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
  "UserID": "4d7c52eb-7d11-4826-9c97-6bea4fb2ae62",
  "transactionType": "send",
  "amount": 100.5,
  "description": "Deposit to savings account",
  "timestamp": "2023-09-20T15:30:00Z"
    }
}

```

### Customer API

### GET Profile
Request :

- Method : GET
- Endpoint : `/customers/profile`
- Header :
  - Accept : application/json
  - Authorization: Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
  "code": "int",
  "status":"string",
  "message": "string",
  "data": {
  "id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
  "userName": "john_doe",
  "email": "johndoe@example.com",
  "password": "securepassword",
  "firstName": "John",
  "lastName": "Doe",
  "bankCard": {
  "id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
  "userID": "4d7c52eb-7d11-4826-9c97-6bea4fb2ae62",
  "bankName": "bca",
  "cardNumber": "123456",
  "cardholderName": "John Doe",
  "expirationDate": "2024-09-30T00:00:00Z",
  "balance": 1000.0,
  "cvv": "123"
}
  },
  "lastLogin": "2023-09-20T15:30:00Z"
}
```

