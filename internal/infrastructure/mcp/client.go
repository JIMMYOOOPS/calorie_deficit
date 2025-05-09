package mcp

import (
	"calorie_deficit/internal/pkg/logger"
	"context"
	"os"

	mcpGolang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

// TODO: Move to transport layer
func createTransport() *stdio.StdioServerTransport {
	// Create a new StdioServerTransport
	clientTransport := stdio.NewStdioServerTransportWithIO(os.Stdin, os.Stdout)
	return clientTransport
}

// Initialize the MCP client
func InitializeClient() {
	// TODO: take transport as an argument
	transport := createTransport()

	// Create a new client
	client := mcpGolang.NewClient(transport)
	logger.Logger.Info("Client created")
	// Initialize the client
	response, err := client.Initialize(context.Background())

	if err != nil {
		logger.Logger.Fatalf("Failed to initialize client: %v", err)
	}
	// Check if the client is initialized
	if response == nil {
		logger.Logger.Info("Client is not initialized")
		return
	}
	logger.Logger.Info("Client is initialized")

	// List available tools
	tools, err := client.ListTools(context.Background(), nil)
	if err != nil {
		logger.Logger.Fatalf("Failed to list tools: %v", err)
	}
	// Print the available tools
	logger.Logger.Info("Available tools:")
	for _, tool := range tools.Tools {
		desc := ""
		if tool.Description != nil {
			desc = *tool.Description
		}
		logger.Logger.Infof("Tool: %s, Description: %s", tool.Name, desc)
	}

	// Print the response
	logger.Logger.Info("Client initialized successfully")
}
