🏦 Omar Bank Backend:
A robust, production-grade backend banking API built from scratch in Go, designed with clean architecture principles, secure authentication, and strict transaction management.

🚀 Key Features:
RESTful API Architecture: Built with the high-performance Gin framework featuring request validation, middleware routing, and clean JSON payload handling.

Type-Safe Database Operations: Utilizes SQLC to generate type-safe Go code from raw SQL queries, eliminating runtime SQL mapping errors.

ACID-Compliant Transactions: Implements complex concurrent-safe transaction logic (TransferTx) that handles account transfers, entry logging, and balance updates atomically within database transactions.

Secure Authentication: Token-based authorization middleware (JWT/PASETO) protecting endpoints for user management, account creation, and fund transfers.

Comprehensive Testing Suite: Fully unit-tested API handlers and database layers using GoMock for dependency mocking, along with stretchr/testify for assertions.

Database Migrations: Managed schema versions smoothly using golang-migrate.

🛠️ Tech Stack
Language: Go

Database: PostgreSQL

Framework & Routing: Gin Web Framework

SQL Generation: SQLC

Testing & Mocking: GoMock, Testify

Authentication: JWT / PASETO