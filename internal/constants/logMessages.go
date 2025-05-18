package constants

type GeneralLogMessages struct {
	InvalidRequest string
}

type DatabaseLogMessages struct {
	PostgresConnectionSuccess string
	PostgresConnectionFailed  string
	PostgresMigration         struct {
		MigrationSuccess string
		MigrationFailed  string
	}
}

type EnvironmentLogMessages struct {
	FileLoaded     string
	FileLoadFailed string
}

type MCPLogMessages struct {
	Client struct {
		ClientInitialized     string
		ClientInitFailed      string
		InvalidModelSpecified string
	}
}

var LogMessages = struct {
	General     GeneralLogMessages
	Database    DatabaseLogMessages
	Environment EnvironmentLogMessages
	MCP         MCPLogMessages
}{
	General: GeneralLogMessages{
		InvalidRequest: "Invalid request",
	},
	Database: DatabaseLogMessages{
		PostgresConnectionSuccess: "Connected to PostgreSQL database successfully",
		PostgresConnectionFailed:  "Failed to connect to PostgreSQL database",
		PostgresMigration: struct {
			MigrationSuccess string
			MigrationFailed  string
		}{
			MigrationSuccess: "Postgres migration successful",
			MigrationFailed:  "Postgres migration failed",
		},
	},
	Environment: EnvironmentLogMessages{
		FileLoaded:     "Environment variables loaded successfully",
		FileLoadFailed: "Error loading .env file",
	},
	MCP: MCPLogMessages{
		Client: struct {
			ClientInitialized     string
			ClientInitFailed      string
			InvalidModelSpecified string
		}{
			ClientInitialized:     "Client: %s, initialized successfully",
			ClientInitFailed:      "Failed to initialize Client: %s",
			InvalidModelSpecified: "Invalid model specified: %s",
		},
	},
}
