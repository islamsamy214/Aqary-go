# Aqary-go

# Setup

## Install dependencies

```bash
go mod tidy
```

## Configure database
    
```bash
cp .env.example .env
```

and update the `.env` file with your database credentials.

## Run the server

```bash
go run main.go
```

# APIs

## Users

### Index

```http
GET /api/users
```

### Create

```http
POST /api/users
```

### Generate OTP

```http
POST /api/users/generateotp
```

### Verify OTP

```http
POST /api/users/verifyotp
```
