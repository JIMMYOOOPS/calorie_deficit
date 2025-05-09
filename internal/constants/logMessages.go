package constants

// Log messages for database operations
const (
	// Log message for successful database connection
	PostgresConnectionSuccess = "Connected to PostgreSQL database successfully"
	// Log message for failed database connection
	PostgresConnectionFailed = "Failed to connect to PostgreSQL database"
)

// Log messages for environment variable loading
const (
	// Log message for successful loading of environment variables
	EnvFileLoaded = "Environment variables loaded successfully"
	// Log message for failed loading of environment variables
	EnvFileLoadFailed = "Error loading .env file"
)
