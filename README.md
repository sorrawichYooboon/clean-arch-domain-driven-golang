# Clean Architecture with Domain-Driven Design in Golang

## Project About

This project is a RESTful API built using Golang, following the principles of Clean Architecture and Domain-Driven Design. It provides endpoints for managing books and authors in a bookstore application. The API supports basic CRUD operations and utilizes GORM for database interactions, Redis for caching, and Echo for handling HTTP requests. Swagger is integrated to provide interactive API documentation.

## How to Set Up

### Prerequisites

- Go (version 1.17 or higher)
- PostgreSQL (or your preferred database)
- Redis (for caching)
- Docker (optional, for running database and Redis in containers)
- Git

### 1. Clone the Repository

```bash
git clone https://github.com/sorrawichYooboon/clean-arch-domain-driven-golang.git
cd clean-arch-domain-driven-golang
```

### 2. Configure Environment Variables

Create a .env file in the root directory and configure the required environment variables.

```bash
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=bookstore
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
```

### 3. Run Docker Compose

To set up the PostgreSQL and Redis services, run the following command:

```base
docker-compose up -d
```

### 4. Check Docker Compose Run Correctly

```base
docker-compose ps
```

## How to Run the Server

### Install Dependencies

Navigate to the root directory and run the following command to install the required Go packages:

```bash
go mod tidy
```

### Start the Server

Run the following command to start the server:

```bash
go run main.go
```

The server will start on port <b>8080</b> by default. You can change the port in the main.go file if needed.

## How to Open Swagger

Once the server is running, you can access the Swagger UI to view the API documentation. [Click here to open Swagger](http://localhost:8080/swagger/index.html) or

```bash
http://localhost:8080/swagger/index.html
```

This will provide an interactive interface for testing the API endpoints.

## Detail Explanation

### Project Structure

The project follows a modular structure adhering to Clean Architecture principles. Here’s a brief overview of the main directories:

- **config/**: Contains the configuration setup for database and Redis connections. It manages environment variables and configuration logic to ensure the application can connect to necessary services seamlessly.

- **internal/**: This directory contains the core application logic, structured as follows:
  - **domain/**: Core domain entities, including definitions for the `Book` and `Author` entities.
  - **usecase/**: This folder implements the use case logic for the application, encapsulating the business rules and application services related to books and authors.
  - **repository/**: Defines repository interfaces that outline how data will be accessed. This layer abstracts the underlying data source, allowing for easy swapping between different implementations.
  - **infrastructure/**: Contains adapters for various infrastructures:
    - **database/**: Implements PostgreSQL-specific repository logic.
    - **cache/**: Implements Redis caching for optimized data retrieval.
    - **http/**: Manages HTTP routing and controllers, handling incoming requests and responses.
- **migrations/**: Contains SQL scripts for creating the required database tables, which are automatically executed by the PostgreSQL Docker container during initialization. This ensures the database is set up correctly for application use.

- **docs/**: Holds the generated Swagger documentation files, providing a user-friendly interface for exploring API endpoints and their functionalities.

### Key Features

- **CRUD Operations**: Create, Read, Update, and Delete functionality for books and authors.
- **Caching**: Uses Redis to cache book and author data for faster access.
- **Swagger Integration**: Automatically generated API documentation for easy - exploration of endpoints.
- **GORM**: Utilizes GORM for ORM functionality with PostgreSQL.

## Conclusion

This project demonstrates the application of Clean Architecture and Domain-Driven Design in building a robust API with Go. Feel free to explore, extend, and adapt the code to fit your needs!
