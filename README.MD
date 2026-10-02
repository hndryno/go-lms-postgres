# LMS - Loan Management System

Backend API untuk Loan Management System menggunakan Go dan PostgreSQL.

## Tech Stack

- Go 1.23+
- Fiber v3
- PostgreSQL
- pgx
- sqlc
- goose

## Features

- Customer CRUD
- Loan Management
- Loan Calculation
- Repayment
- Pagination
- PostgreSQL Stored Procedure

## Project Structure

```text
cmd/
└── api/

internal/
├── customer/
├── loan/
├── repayment/
├── config/
├── response/
└── adapters/
    └── postgresql/
        ├── db/
        ├── migrations/
        └── queries/
```

Project menggunakan **feature-oriented structure**, sehingga logic yang berkaitan dengan setiap feature dikelompokkan dalam satu folder.

## Installation

Clone repository:

```bash
git clone <repository-url>
cd go-lms-postgresql
```

Install dependencies:

```bash
go mod download
```

Copy `.env.example` menjadi `.env` dan sesuaikan konfigurasi PostgreSQL.

## Database

Create database:

```text
lms_db
```

Run migration:

```bash
make migrate-up
```

Generate sqlc:

```bash
make sqlc
```

## Run

```bash
go run ./cmd/api
```

atau menggunakan Air:

```bash
make dev
```

API:

```text
http://localhost:8080
```

## API Endpoints

### Customer

```text
POST   /api/v1/customers/
GET    /api/v1/customers/
GET    /api/v1/customers/{id}
PUT    /api/v1/customers/{id}
DELETE /api/v1/customers/{id}
```

### Loan

```text
POST   /api/v1/loans/
GET    /api/v1/loans/
GET    /api/v1/loans/{id}
```

### Repayment

```text
POST   /api/v1/repayments/
```

## Loan Calculation

```text
Interest Amount = Principal Amount × Interest Rate / 100
Total Amount    = Principal Amount + Interest Amount
```

## PostgreSQL Stored Procedure

Stored procedure digunakan untuk proses:

- Create Loan
- Create Repayment

Script tersedia di:

```text
internal/adapters/postgresql/migrations/
```

## Testing

API telah diuji menggunakan Postman.

Postman collection dan screenshot hasil testing tersedia pada file submission.

## Author

**Hendriyono**

GitHub: https://github.com/hndryno  
Email: hendriyono97@gmail.com