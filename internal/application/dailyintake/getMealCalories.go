// interal/application/dailyintake/getMealCalories.go
package dailyintakeapplication

import (
	"calorie_deficit/internal/constants"
	dailyintakedto "calorie_deficit/internal/dto/dailyintake"
	"calorie_deficit/internal/infrastructure/mcp"
	"calorie_deficit/internal/pkg/logger"
	"calorie_deficit/internal/types"

	"context"
	"encoding/json"
)

type MCPDailyIntakeService struct {
	MCPClient mcp.LLMClient
}

func NewMCPDailyIntakeService(mcpClient mcp.LLMClient) *MCPDailyIntakeService {
	return &MCPDailyIntakeService{
		MCPClient: mcpClient,
	}
}

// init the types to be used imported from dailyintake package
type DailyIntakeCreateRequestDTO = dailyintakedto.DailyIntakeCreateRequestDTO
type MealItemRequestDTO = dailyintakedto.MealItemRequestDTO
type DailyIntakeCreateServiceDTO = dailyintakedto.DailyIntakeCreateServiceDTO
type MealItemDTO = dailyintakedto.MealItemDTO

var (
	// ErrEmptyResponse is an error for empty response from the LLM client
	errEmptyResponse = types.AppError{
		Code:    400,
		Message: constants.LogMessages.MCP.Client.EmptyResponse,
	}
)

func (s *MCPDailyIntakeService) GetMealCalories(params DailyIntakeCreateRequestDTO, mealItems []MealItemRequestDTO) (DailyIntakeCreateServiceDTO, error) {
	// Call the LLM client to get the meal calories for meal items
	mealItemsDTO := make([]MealItemRequestDTO, len(mealItems))
	for i, item := range mealItems {
		mealItemsDTO[i] = MealItemRequestDTO(item)
	}
	// Form the prompt for the LLM client
	jsonMealItemsDTO, err := json.Marshal(mealItemsDTO)
	if err != nil {
		return DailyIntakeCreateServiceDTO{}, err
	}
	// The prompt should be a string that describes the meal items and asks for their calorie count
	userRoleInput := "Given the following meal items, estimate the calorie count for each and return ONLY a JSON array with all fields filled in (name, measurement, quantity, calorie). Do not include any explanation or extra text.\n\n" + string(jsonMealItemsDTO)
	schema := map[string]interface{}{
		"type": "array",
		"items": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "Name of the meal item",
				},
				"measurement": map[string]interface{}{
					"type":        "string",
					"description": "the measurement of the meal item should be one of the following: grams, ml",
				},
				"calorie": map[string]interface{}{
					"type":        "number",
					"description": "the calorie count of the meal item",
				},
				"quantity": map[string]interface{}{
					"type":        "number",
					"description": "the quantity of the meal item",
				},
			},
			"required":             []string{"name", "measurement", "calorie", "quantity"},
			"additionalProperties": false,
		},
		"description": "A list of meal items with their calorie counts",
	}

	promptResponse, err := s.MCPClient.CreateChatCompletion(context.Background(), userRoleInput, schema)
	if err != nil {
		logger.Logger.Errorf("Error calling LLM client: %v", err)
		return DailyIntakeCreateServiceDTO{}, err
	}
	if promptResponse == "" {
		logger.Logger.Error("Empty response from LLM client, Please try again")
		return DailyIntakeCreateServiceDTO{}, &errEmptyResponse
	}
	logger.Logger.Infof("Raw LLM response: '%s'", promptResponse)
	// Parse the response from the LLM client
	var mealCaloriesResponse []MealItemDTO
	err = json.Unmarshal([]byte(promptResponse), &mealCaloriesResponse)
	if err != nil {
		logger.Logger.Errorf("Error unmarshalling LLM response: %v", err)
		return DailyIntakeCreateServiceDTO{}, err
	}
	logger.Logger.Infof("Meal calories response: %v", mealCaloriesResponse)
	// Map the response to the DailyIntakeCreateServiceDTO
	mealCaloriesResponseDTO := DailyIntakeCreateServiceDTO{
		UserID:    params.UserID,
		Date:      params.Date,
		MealType:  params.MealType,
		MealItems: mealCaloriesResponse,
	}
	logger.Logger.Infof("Meal calories response DTO: %v", mealCaloriesResponseDTO)
	return mealCaloriesResponseDTO, nil
}
