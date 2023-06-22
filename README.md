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
ketika username atau pun email kamu sudah pernah digunakan akan mengembalikan response error

## Response
```
{
  "code": 500,
    "message": "Error 1062 (23000): Duplicate entry 'achyarbagus' for key 'user_name'"
}
```

## Sign in

kamu bisa melakukan login dengan menggunakan api url POST:  "/api/v1/sign-in"

## Response
```
{
    "username":"achyarbagus",
    "password":"123"
}
```
ketika username kamu belum terdaftar maka akan mengembalikan response error

## Response
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

### Example

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
