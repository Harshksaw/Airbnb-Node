# Airbnb-Node Microservices Architecture

This repository contains a microservices-based architecture for an Airbnb-like application, built with multiple services handling different aspects of the system. The project is designed for scalability, maintainability, and ease of deployment using Docker.

## Project Overview

The system is divided into several microservices, each responsible for a specific domain:

- **Authentication Service (AuthInGo)**: Handles user authentication, roles, and permissions using Go.
- **Booking Service**: Manages booking operations using TypeScript/Node.js.
- **Hotel Service**: Handles hotel and room management using TypeScript/Node.js.
- **Notification Service**: Manages notifications and messaging using TypeScript/Node.js.

Additional components include Docker setup for containerization and MySQL initialization scripts.

## Services

### AuthInGo (Go)

**Location:** `AuthInGo/`

**Description:** The authentication service manages user registration, login, JWT token generation, and role-based access control (RBAC). It includes user management, role assignment, and permission handling.

**Features:**

- User registration and login with password hashing
- JWT-based authentication
- Role and permission management
- User-role assignments
- Secure password handling with bcrypt
- Repository pattern for data access

**Technologies:** Go, JWT, MySQL, Docker

**Key Components:**

- `services/user_service.go`: User management logic
- `services/role_service.go`: Role and permission logic
- `controllers/`: HTTP handlers
- `models/`: Data models
- `repositories/`: Data access layer

### BookingService (TypeScript/Node.js)

**Location:** `BookingService/`

**Description:** The booking service handles all booking-related operations, including creating, updating, and managing reservations for properties.

**Features:**

- Booking creation and management
- Integration with Prisma ORM for database operations
- Validation using custom validators
- Middleware for correlation and error handling
- Queue-based processing for asynchronous tasks

**Technologies:** Node.js, TypeScript, Prisma, Redis, Docker

**Key Components:**

- `services/booking.service.ts`: Core booking logic
- `controllers/booking.controller.ts`: API endpoints
- `repositories/booking.repository.ts`: Data access
- `prisma/`: Database schema and migrations

### HotelService (TypeScript/Node.js)

**Location:** `HotelService/`

**Description:** The hotel service manages hotel listings, room categories, and related operations.

**Features:**

- Hotel and room management
- Room generation and publishing to queues
- Producer-consumer pattern for message handling
- Integration with Redis for caching
- Comprehensive validation and error handling

**Technologies:** Node.js, TypeScript, Redis, Docker

**Key Components:**

- `services/hotel.service.ts`: Hotel management logic
- `services/roomGeneration.service.ts`: Room generation logic
- `producers/roomGeneration.producer.ts`: Message publishing
- `controllers/`: API handlers
- `repositories/`: Data access layer

### NotificationService (TypeScript/Node.js)

**Location:** `NotificationService/`

**Description:** The notification service handles sending notifications, emails, and other messaging to users.

**Features:**

- Email sending via mailer service
- Template-based notifications
- Queue processing for reliable delivery
- Integration with external services

**Technologies:** Node.js, TypeScript, Redis, Docker

**Key Components:**

- `services/mailer.service.ts`: Email sending logic
- `templates/`: Email templates
- `processors/`: Message processing
- `producers/`: Message publishing

## Infrastructure

### Docker

**Location:** `Docker/`

Contains Docker Compose files for orchestrating all services, databases, and dependencies.

### MySQL Initialization

**Location:** `mysql-init/`

Scripts for initializing the MySQL database with schemas, users, and initial data.

## Getting Started

1. **Prerequisites:**
   - Docker and Docker Compose
   - Node.js (for TypeScript services)
   - Go (for AuthInGo)

2. **Clone the Repository:**

   ```bash
   git clone https://github.com/Harshksaw/Airbnb-Node.git
   cd Airbnb-Node
   ```

3. **Start Services:**

   ```bash
   cd Docker
   docker-compose up -d
   ```

4. **Individual Service Setup:**
   - For each service, navigate to its directory and follow the README instructions.

## Architecture

The system follows a microservices architecture with:

- RESTful APIs for inter-service communication
- Message queues (Redis) for asynchronous processing
- Shared database (MySQL) with service-specific schemas
- Containerized deployment with Docker

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes and test
4. Submit a pull request

## License

MIT License

## Contact

For questions or issues, please open a GitHub issue or contact the maintainers.