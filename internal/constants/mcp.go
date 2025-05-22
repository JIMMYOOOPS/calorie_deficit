// internal/constants/mcp.go

package constants

type LLM struct {
	Name   string
	Models struct {
		Google []string
		Meta   []string
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
			Meta   []string
		}{
			Google: []string{}, // select models that support Structured Output response
			Meta:   []string{"meta-llama/llama-4-maverick:free", "meta-llama/llama-3.3-8b-instruct:free"},
		},
	},
}
