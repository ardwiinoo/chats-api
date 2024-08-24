# API SPEC

## User

### Register

Endpoint: POST /api/users/signup

Request Body :

```json
{
    "username": "user",
    "email": "user@gmail.com",
    "password": "rahasia"
}
```

Response Body (Success) :

```json
{
    "id": 1,
    "username": "user",
    "email": "user@gmail.com"
}
```

Response Body (Failed) :

```json
{
    "error": "EOF"
}
```
