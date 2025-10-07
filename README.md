# Ecommerce Site in Go

This project is an implementation of a basic Ecommerce site using Go. It features two main domains: **User** and **Product**.

## Features

### User Domain
- **User Creation:** Register new users.
- **User List View:** View all registered users.
- **JWT Authentication:** Secure endpoints with JSON Web Tokens.

### Product Domain
- **Product Creation:** Add new products.
- **Product List View:** View all available products.
- **Product Update:** Edit product details.

## Getting Started

1. **Clone the repository**
    ```bash
    cd GoLang
    ```

2. **Install dependencies**
    ```bash
    go mod tidy
    ```

3. **Run the application**
    ```bash
    go run main.go
    ```
## API Endpoints

    ### Product Endpoints

    - `POST /products` — Create a new product
    - `GET /products` — List all products
    - `GET /products/{id}` — Get details of a specific product
    - `PUT /products/{id}` — Replace a product's details
    - `PATCH /products/{id}` — Update part of a product's details
    - `DELETE /products/{id}` — Remove a product

    ### User Endpoints

    - `GET /users` — List all users
    - `GET /users/{id}` — Get details of a specific user
    - `POST /users` — Register a new user
    - `POST /users/login` — Authenticate a user and receive a JWT
    - `PUT /users/{id}` — Replace a user's details
    - `PATCH /users/{id}` — Update part of a user's details
    - `DELETE /users/{id}` — Remove a user

## Technologies Used

- Go (Golang)
- JWT for authentication

## License

This project is licensed under the MIT License.