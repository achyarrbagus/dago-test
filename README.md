# Dago Test

Web Api Aplication with golang echo and mysql.

## Daftar Isi

- [Instalasi](#instalasi)


## Instalasi

Langkah-langkah untuk menginstal proyek ini:

1. Clone repositori ini: `https://github.com/achyarrbagus/dago-test.git`
2. Masuk ke direktori proyek: `cd dago-test`
3. Instal dependensi: `go mod init`

## Penggunaan

Cara menggunakan proyek ini:

1. Jalankan aplikasi: `npm start`
2. konnfigurasi file .env
```
DB_HOST = localhost | your host
DB_PORT = 5432 | your Db port
DB_USER = root | your Db user
DB_PASSWORD = 123 | your Db password
DB_NAME = dago-test | your Db name
PORT = 5000 | your Aplication port
```
4. Buka aplikasi di browser dengan URL: `http://localhost:5000`


## Sign Up 

kamu bisa melakukan pendaftaran dengan menggunakan api url POST:  "/api/v1/sign-up" 

## Request
```
{
    "username": "achyarrbagus",
    "email": "achyarbagus@gmail.com",
    "password": "123",
    "phone": "123",
    "Address": "bekasi"
}

```
## Response

- **Code**: 200 OK
- **Content-Type**: application/json
```
{
    "code": 200,
    "data": {
        "id": 6,
        "username": "achyarbagus",
        "email": "achyarrbagus@gmail.com",
        "phone": "123",
        "password": "$2a$10$mwNVE6ELVjrnUeUfJxpX7.mCxwg1xbO2Jzfh3rqQVfcmbc9XUdY7C",
        "address": "bekasi"
    }
}
```

ketika username atau pun email kamu sudah pernah digunakan akan mengembalikan response error
## Response
- **Code**: 500
- **Content-Type**: application/json

```
{
  "code": 500,
    "message": "Error 1062 (23000): Duplicate entry 'achyarbagus' for key 'user_name'"
}
```

## Sign in

kamu bisa melakukan login dengan menggunakan api url POST:  "/api/v1/sign-in"

## Request
```
{
    "username":"achyarbagus",
    "password":"123"
}
```
## Response

- **Code**: 200 OK
- **Content-Type**: application/json

```
{
    "code": 200,
    "data": {
        "username": "achyarbagus",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc0NjE2ODAsImlkIjo2fQ.5L-xOsX5VtasxhgKUbrBEZA4hsbkBxLvKP8D7xtazdM",
        "userId": 6
    }
}
```

ketika username kamu belum terdaftar maka akan mengembalikan response error
- **Code**: 500 
- **Content-Type**: application/json
    
```json
{
    "code": 404,
    "message": "record not found"
}
```
## GetAllUser

kamu bisa mendapat kan semua data user yang terdaftar dengan menggunakan api url GET: "/api/v1/users"

## Response

- **Code**: 200 OK
- **Content-Type**: application/json

```json
{
    "code": 200,
    "data": [
        {
            "id": 1,
            "username": "baga",
            "email": "ass@gmail.com",
            "phone": "123",
            "password": "$2a$10$OVd1yDhKhMXpyS2hJwCzQeydoAvfU8O7I5LZ.0ar9AS11gIDCHRlG",
            "address": "bekasi"
        },
        {
            "id": 2,
            "username": "achyarrbagus",
            "email": "achyarbagus@gmail.com",
            "phone": "123",
            "password": "$2a$10$5zsMNXy0pIXlq261jovkk.QRnQxy9c3Ru17B7VFM6w.A9xwKr1ewy",
            "address": "bekasi"
        }
    ]
}
```

## GetUserById

kamu bisa mendapat kan semua data user yang terdaftar berdasarkan id dengan menggunakan api url GET: "/api/v1/user/:id"

## Response

- **Code**: 200 OK
- **Content-Type**: application/json

```json
{
    "code": 200,
    "data": {
        "id": 1,
        "username": "baga",
        "email": "ass@gmail.com",
        "phone": "123",
        "password": "$2a$10$OVd1yDhKhMXpyS2hJwCzQeydoAvfU8O7I5LZ.0ar9AS11gIDCHRlG",
        "address": "bekasi"
    }
}

```
ketika data user kamu belum terdaftar maka akan mengembalikan response error

## Response

- **Code**: 404 
- **Content-Type**: application/json
    
```json
{
    "code": 404,
    "message": "record not found"
}
```


