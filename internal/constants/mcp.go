package constants

type LLM struct {
	Name string
}

// LLMs contains the list of available LLMs
var LLMs = struct {
	OpenAI       LLM
	OpenRouterAI LLM
}{
	OpenAI: LLM{
		Name: "openai",
	},
	OpenRouterAI: LLM{
		Name: "openRouter",
	},
}
