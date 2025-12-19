# Content Platform API

A scalable backend API for managing, publishing, and scheduling content.  
This project is designed to simulate a real-world production backend, focusing on clean architecture, scalability, and best practices commonly used in modern software companies.

---

## Project Overview

The **Content Platform API** provides a secure and extensible backend for content-driven platforms such as blogs, corporate websites, and SaaS products.

It handles:
- User authentication and authorization
- Content creation and retrieval
- Secure access control
- Future scalability for background jobs, caching, and integrations

This project was built with a strong emphasis on backend engineering and system design rather than UI complexity.

---

## Architecture

The project follows a **layered architecture**, separating responsibilities to improve maintainability and scalability.

Client (Web / Mobile / API Consumer)
|
v
HTTP Handlers
|
v
Services
|
v
Repositories
|
v
PostgreSQL

yaml
Copiar código

### Layer responsibilities

- **Handlers**: Handle HTTP requests and responses
- **Services**: Contain business logic and validation
- **Repositories**: Handle database access
- **Models**: Define core domain structures

---

## Tech Stack

- Go — Backend API
- PostgreSQL — Relational database
- Docker & Docker Compose — Local environment and containerization
- JWT — Authentication and authorization
- SQL — Data persistence
- REST API — Communication protocol

---

## Features

### Implemented
- User registration
- User authentication (JWT)
- Create content posts
- List published posts
- Health check endpoint

### Planned
- Post scheduling
- Background workers
- Caching with Redis
- Rate limiting
- Role-based access control
- API documentation (OpenAPI/Swagger)
- Integration with a Node.js BFF

---

## Project Structure

content-platform-api/
├── cmd/
│ └── api/
│ └── main.go # Application entry point
├── internal/
│ ├── handler/ # HTTP handlers
│ ├── service/ # Business logic
│ ├── repository/ # Database access
│ ├── model/ # Domain models
│ └── auth/ # Authentication logic
├── migrations/ # Database migrations
├── pkg/ # Shared utilities
├── docker-compose.yml # Local environment
├── Dockerfile # API container
├── go.mod
└── README.md

yaml
Copiar código

---

## Authentication

The API uses **JWT (JSON Web Tokens)** for authentication.

- Tokens are issued upon successful login
- Protected endpoints require a valid JWT in the `Authorization` header

Example:

Authorization: Bearer <token>

yaml
Copiar código

---

## API Endpoints (Initial)

### Health Check
GET /health

shell
Copiar código

### Authentication
POST /auth/register
POST /auth/login

shell
Copiar código

### Posts
POST /posts (authenticated)
GET /posts (public)

yaml
Copiar código

---

## Running the Project Locally

### Prerequisites

- Go 1.22+
- Docker
- Docker Compose

### Clone the repository

git clone https://github.com/your-username/content-platform-api.git
cd content-platform-api

shell
Copiar código

### Run with Docker

docker-compose up --build

arduino
Copiar código

The API will be available at:

http://localhost:8080

yaml
Copiar código

---

## Development Principles

- Clean and readable code
- Clear separation of concerns
- Meaningful commit messages
- English-only documentation and commits
- Focus on real-world backend scenarios

---

## Future Improvements

- Introduce background jobs using message queues
- Implement caching strategies
- Add monitoring and structured logging
- Improve security and performance
- Expand test coverage

---

## Contribution

This project is currently for learning and portfolio purposes.  
Feedback, suggestions, and code reviews are always welcome.

---

## License

This project is licensed under the MIT License.

---

## Contact

**Author:** Guilherme Souza  
**Role:** Backend Engineer (Go / Node.js)  
**Location:** Brazil  
**GitHub:** https://github.com/GuilhermeFerza
