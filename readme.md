# Vivo Project

Vivo is a Go-based application that utilizes the Fiber web framework and JWT for authentication.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Running the Application](#running-the-application)
4. [Configuration](#configuration)

## Prerequisites

- Go (version 1.18 or later)
- A PostgreSQL database

## Installation

1. Clone the repository:

```
git clone https://github.com/yourusername/vivo.git
cd vivo
```

2. Install dependencies using Go modules:

```
go mod tidy
```

## Configuration

Create a .env file in the root directory of your project with the following fields:
```
POSTGRES_HOST=your_host
POSTGRES_PORT=your_port
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_database

JWT_SECRET=your_jwt_secret
JWT_EXPIRED_IN=your_jwt_expiration_time
JWT_MAXAGE=your_jwt_max_age
```



## Running the Application

1. Ensure that your database is set up and the connection details are correctly configured in your environment variables or configuration files.

2. Run the application:
```
go run main.go
```

3. The application will start on http://localhost:8000.



