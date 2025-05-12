package constants

type DatabaseLogMessages struct {
	PostgresConnectionSuccess string
	PostgresConnectionFailed  string
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
	Database    DatabaseLogMessages
	Environment EnvironmentLogMessages
	MCP         MCPLogMessages
}{
	Database: DatabaseLogMessages{
		PostgresConnectionSuccess: "Connected to PostgreSQL database successfully",
		PostgresConnectionFailed:  "Failed to connect to PostgreSQL database",
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
