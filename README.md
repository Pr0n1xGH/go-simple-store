# Go-Start REST API

A RESTful API for user management and shopping cart operations.

## Base URL

```
http://localhost:8081
```

---

## Authentication

- **JWT** is used for authentication.
- Obtain tokens via `/api/login`.
- For protected endpoints, send the `Authorization: Bearer <access_token>` header.
- If the access token is expired, send `X-Refresh-Token: <refresh_token>` header to refresh.

---

## Endpoints

### Auth

#### Login

**POST** `/api/login`

Authenticate user and receive JWT tokens.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

**Response:**
```json
{
  "access_token": "string",
  "refresh_token": "string"
}
```

---

#### Refresh Token

**POST** `/api/refresh`

Get new tokens using a refresh token.

**Request Body:**
```json
{
  "refresh_token": "string"
}
```

**Response:**
```json
{
  "access_token": "string",
  "refresh_token": "string"
}
```

---

### Users

> **All user endpoints require JWT authentication.**

#### Create User

**POST** `/api/user`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "yourpassword"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "password": "yourpassword"
}
```

---

#### Get All Users

**GET** `/api/user`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "password": "yourpassword"
  }
]
```

---

#### Get User by ID

**GET** `/api/user/{id}`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "password": "yourpassword"
}
```

---

#### Update User

**PUT** `/api/user/{id}`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "password": "newpassword"
}
```

**Response:**
```json
{
  "message": "Пользователь обновлён"
}
```

---

#### Delete User

**DELETE** `/api/user/{id}`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Response:**
```json
{
  "message": "Пользователь удалён"
}
```

---

### Cart

> **All cart endpoints require JWT authentication.**

#### Create Cart

**POST** `/api/cart/create`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "user_id": 1
}
```

**Response:**
```json
{
  "ID": 1,
  "CreatedAt": "2024-06-13T12:00:00Z",
  "UpdatedAt": "2024-06-13T12:00:00Z",
  "DeletedAt": null,
  "user_id": 1
}
```

---

### Cart Items

> **All cart item endpoints require JWT authentication.**

#### Add Item to Cart

**POST** `/api/cart/{cartID}/items`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "product": "Product Name",
  "quantity": 2
}
```

**Response:**
```json
{
  "ID": 1,
  "CartID": 1,
  "Product": "Product Name",
  "Quantity": 2,
  "Cart": { ... }
}
```

---

#### Get Cart Items

**GET** `/api/cart/{cartID}/items`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Response:**
```json
[
  {
    "ID": 1,
    "CartID": 1,
    "Product": "Product Name",
    "Quantity": 2,
    "Cart": { ... }
  }
]
```

---

#### Update Cart Item Quantity

**PUT** `/api/cart/items/{itemID}`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Request Body:**
```json
{
  "quantity": 5
}
```

**Response:**
```json
{
  "message": "Количество обновлено"
}
```

---

#### Delete Cart Item

**DELETE** `/api/cart/items/{itemID}`

**Headers:**
```
Authorization: Bearer <access_token>
```

**Response:**
```json
{
  "message": "Товар удалён из корзины"
}
```

---

## Error Handling

- All errors are returned as plain text or JSON with appropriate HTTP status codes.
- Common status codes: `400 Bad Request`, `401 Unauthorized`, `404 Not Found`, `500 Internal Server Error`.

---

## Notes

- All request/response bodies are in JSON.
- JWT tokens are required for all endpoints except `/api/login` and `/api/refresh`.
- Use the `refresh_token` to obtain a new `access_token` when expired.

---

## Example Usage

### Login and Use JWT

1. **Login:**
   - POST `/api/login` with email and password.
   - Receive `access_token` and `refresh_token`.

2. **Authenticated Request:**
   - Add header: `Authorization: Bearer <access_token>` to all protected endpoints.

3. **Refresh Token:**
   - POST `/api/refresh` with `refresh_token` to get new tokens.

---

## Environment Variables

Set in `.env`:

```
ACCESS_SECRET=your_access_secret
REFRESH_SECRET=your_refresh_secret
```

---

## License

MIT
