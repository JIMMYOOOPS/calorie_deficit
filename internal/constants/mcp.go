// internal/constants/mcp.go

package constants

type LLM struct {
	Name   string
	Models struct {
		Google []string
	}
}

// LLMs contains the list of available LLMs
var LLMs = struct {
	OpenAI     LLM
	OpenRouter LLM
}{
	OpenAI: LLM{
		Name: "openai",
	},
	OpenRouter: LLM{
		Name: "openRouter",
		Models: struct {
			Google []string
		}{
			Google: []string{"google/gemma-3n-e4b-it:free"},
		},
	},
}
