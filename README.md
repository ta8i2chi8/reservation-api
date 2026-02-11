# Reservation System API

A clean architecture + DDD reservation system built with Go, implementing JWT authentication and CRUD operations.

## Architecture

This project follows Clean Architecture principles with Domain-Driven Design (DDD):

- **Domain Layer**: Business entities and logic
- **Use Case Layer**: Application business rules
- **Interface Layer**: HTTP handlers and middleware
- **Infrastructure Layer**: Database, JWT, and external services

## Features

- JWT-based authentication and authorization
- User management (CRUD)
- Reservation system with time slots and capacity management
- RESTful API with custom HTTP router
- PostgreSQL database with GORM
- Docker containerization
- CI/CD with GitHub Actions

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for local development)

### Using Docker Compose (Recommended)

1. Clone the repository
2. Copy environment file:
   ```bash
   cp .env.example .env
   ```
3. Start the services:
   ```bash
   make dev
   ```

The API will be available at `http://localhost:8080` and the database at `localhost:5432`.

### Local Development

1. Set up PostgreSQL database
2. Configure environment variables
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the application:
   ```bash
   make run
   ```

## API Endpoints

### Authentication

- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login user
- `POST /api/auth/validate` - Validate JWT token

### Users

- `POST /api/users` - Create user
- `GET /api/users?id={id}` - Get user by ID
- `POST /api/users/login` - Login (alternative endpoint)

### Reservations

- `POST /api/reservations` - Create reservation (requires auth)
- `GET /api/reservations?id={id}` - Get reservation by ID (requires auth)
- `GET /api/reservations/user?user_id={id}` - Get user reservations (requires auth)
- `POST /api/reservations/confirm` - Confirm reservation (requires auth)
- `DELETE /api/reservations?reservation_id={id}&user_id={id}` - Cancel reservation (requires auth)

## Database Schema

The system uses a normalized database structure with DDD patterns:

### Users Table
- `id` (PK)
- `email` (unique)
- `password`
- `name`
- `created_at`
- `updated_at`

### Reservations Table
- `id` (PK)
- `user_id` (FK)
- `date` (embedded from TimeSlot)
- `start_time` (embedded from TimeSlot)
- `end_time` (embedded from TimeSlot)
- `capacity` (embedded from TimeSlot)
- `status`
- `created_at`
- `updated_at`

## Testing

Run the test suite:

```bash
make test
```

## Development Commands

```bash
make help        # Show all available commands
make build       # Build the application
make run         # Build and run
make dev         # Start with docker-compose
make test        # Run tests
make clean       # Clean artifacts
make format      # Format code
make lint        # Run linter (requires golangci-lint)
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | localhost | Database host |
| `DB_PORT` | 5432 | Database port |
| `DB_USER` | postgres | Database user |
| `DB_PASSWORD` | password | Database password |
| `DB_NAME` | reservation_system | Database name |
| `DB_SSLMODE` | disable | SSL mode |
| `JWT_SECRET` | - | JWT secret key |
| `PORT` | 8080 | API server port |

## CI/CD

The project includes GitHub Actions workflow that:

1. Runs tests on each push/PR
2. Builds Docker image on main branch
3. Pushes to GitHub Container Registry (GHCR)

## Security Notes

- Change JWT secret in production
- Use proper password hashing (currently simplified for demo)
- Configure proper CORS origins in production
- Enable database SSL in production
- Set up proper database user permissions

## License

This project is for educational purposes to demonstrate clean architecture patterns in Go.