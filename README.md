
# Monitoring Project - UserService Contribution

## Summary
This contribution adds a complete UserService module to the monitoring system.

## Features
- Gin-based API
- MongoDB connection using environment variable
- Modular code (routes, controllers, models)
- Ready to expand with AWS S3 or SNS

## Setup

### Prerequisites
- Go 1.18+
- MongoDB (local or Atlas)
- (Optional) AWS CLI

### Installation

1. Copy `.env` and update with your Mongo URI.
2. Run `go mod tidy`
3. Start server: `go run main.go`

## API Endpoints

- `POST /register`
- `POST /login`
- `GET /user/:id`

# Monitoring Project - UserService Contribution

## Summary
This contribution adds a complete UserService module to the monitoring system.

## Features
- Gin-based API
- MongoDB connection using environment variable
- Modular code (routes, controllers, models)
- Ready to expand with AWS S3 or SNS

## Setup

### Prerequisites
- Go 1.18+
- MongoDB (local or Atlas)
- (Optional) AWS CLI

### Installation

1. Copy `.env` and update with your Mongo URI.
2. Run `go mod tidy`
3. Start server: `go run main.go`

## API Endpoints

- `POST /register`
- `POST /login`
- `GET /user/:id`
