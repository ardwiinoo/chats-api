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

### Login

Endpoint: POST /api/users/signin

Request Body :

```json
{
    "email": "user@gmail.com",
    "password": "rahasia"
}
```

Response Body (Success) :

**Set-Cookie**: jwt=token; HttpOnly; Path=/; Max-Age=3600

```json
{
    "id": 1,
    "username": "user"
}
```

Response Body (Failed) :

```json
{
    "error": "EOF"
}
```

### Logout

Endpoint: GET /api/users/logout

**Set-Cookie**: jwt=token; HttpOnly; Path=/; Max-Age=3600

Response Body (Success) :

```json
{
    "message": "logout successful"
}
```

Response Body (Failed) :

```json
{
    "error": "EOF"
}
```

## Websocket

### Create Room

Endpoint: POST /ws/create-room

Request Body :

```json
{
    "id": 1,
    "name": "room1"
}
```

Response Body (Success) :

```json
{
    "id": 1,
    "name": "room1"
}
```

Response Body (Failed) :

```json
{
    "error": "EOF"
}
```

### Join Room

Endpoint: ws://localhost:8080/ws/join-room/:roomId?userId=1&username=user

Response Body (Connected) :

```text
Connected to ws://localhost:8080/ws/join-room/:roomId?userId=1&username=user

Handshake details
...
Request Method: "GET"
Status Code: "101 Switching Protocols"
...
```

Response Body (Failed) :

```json
{
    "error": "EOF"
}
```

### Get Rooms

Endpoint: GET /ws/get-rooms

Response Body (Success) :

```json
{
    "id": 1,
    "name": "room1"
},
{
    "id": 2,
    "name": "room2"
},
{
    "id": 3,
    "name": "room3"
}
```

### Get Clients

Endpoint: GET /ws/get-clients

Response Body (Success) :

```json
{
    "id": 1,
    "username": "user1"
},
{
    "id": 2,
    "username": "user2"
},
{
    "id": 3,
    "username": "user3"
}
```
