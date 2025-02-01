# Readme



```
wedding-app/
├── cmd/                         # Main application entry point
│   ├── main.go                  # Initializes dependencies and starts the server
│
├── config/                      # Configuration files (env, database config)
│   ├── config.go                 # Loads environment variables
│
├── db/                          # Database-related files
│   ├── sqlc/                    # Auto-generated Go code from sqlc
│   ├── queries.sql               # SQL queries for sqlc
│
├── models/                      # Domain models (independent of sqlc)
│   ├── guest.go                  # Guest domain model
│   ├── event.go                  # Event domain model
│   ├── relationship.go           # Relationship model
│
├── repository/                  # Repository interfaces and implementations
│   ├── guest_repository.go       # Guest repository interface
│   ├── postgres_guest_repository.go  # PostgreSQL implementation
│   ├── memory_guest_repository.go    # In-memory implementation
│   ├── event_repository.go       # Event repository interface
│
├── services/                    # Business logic layer
│   ├── guest_service.go          # Business logic for guests
│   ├── event_service.go          # Business logic for events
│
├── handlers/                    # HTTP handlers for Echo framework
│   ├── guest_handler.go          # API handlers for guests
│   ├── event_handler.go          # API handlers for events
│
├── routes/                      # Echo route definitions
│   ├── routes.go                 # Registers routes
│
├── middlewares/                 # Middleware for authentication, logging
│   ├── auth_middleware.go        # JWT authentication middleware
│
├── storage/                     # Storage-related utilities (S3, file uploads)
│   ├── s3_storage.go             # S3 helper functions
│
├── tests/                       # Unit and integration tests
│   ├── guest_test.go             # Tests for guests
│
├── .env                         # Environment variables
├── sqlc.yaml                    # sqlc configuration
├── go.mod                        # Go module dependencies
├── go.sum                        # Go module checksums
```










```
wedding-app/
├── cmd/                 # Entry points for the application
│   └── main.go          # Main file to start the server
├── configs/             # Configuration files (e.g., environment variables)
│   └── config.yaml      # Example YAML configuration
├── internal/            # Application-specific code (not exposed as packages)
│   ├── auth/            # Authentication logic
│   │   ├── middleware.go
│   │   ├── jwt.go
│   │   └── oauth.go     # For Google OAuth
│   ├── admin/           # Admin-specific features
│   │   ├── handlers.go  # Handlers for admin endpoints
│   │   └── routes.go    # Route definitions
│   ├── guests/          # Guest-related features
│   │   ├── handlers.go  # Handlers for guest endpoints
│   │   ├── models.go    # Guest models and database logic
│   │   └── routes.go    # Route definitions
│   ├── events/          # Event management
│   │   ├── handlers.go  # Handlers for event endpoints
│   │   ├── models.go    # Event models and database logic
│   │   └── routes.go    # Route definitions
│   └── reservations/    # Reservation management
│       ├── handlers.go  # Handlers for reservation endpoints
│       ├── models.go    # Reservation models and database logic
│       └── routes.go    # Route definitions
├── pkg/                 # Shared reusable code (can be imported into other projects)
│   ├── utils/           # Utility functions
│   │   ├── hash.go      # Password hashing utilities
│   │   ├── response.go  # Standardized API responses
│   │   └── validation.go # Input validation logic
│   ├── logger/          # Logging utilities
│   │   └── logger.go
│   └── database/        # Database connection and migrations
│       ├── connection.go
│       └── migrations/
│           ├── 001_create_users_table.sql
│           └── 002_create_guests_table.sql
├── web/                 # Frontend files (if serving templates or a single-page app)
│   ├── templates/       # HTML templates for server-side rendering
│   ├── static/          # Static assets like CSS, JS, images
│   └── dist/            # Built frontend files (if using React/Vue)
├── go.mod               # Go module file
├── go.sum               # Go module dependencies
└── README.md            # Documentation
```

