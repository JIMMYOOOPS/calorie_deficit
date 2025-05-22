// interal/application/dailyintake/getMealCalories.go
package dailyintakeapplication

import (
	dailyintakedto "calorie_deficit/internal/dto/dailyintake"
	"calorie_deficit/internal/infrastructure/mcp"
	"calorie_deficit/internal/pkg/logger"

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
type DailyIntakeCreateServiceDTO = dailyintakedto.DailyIntakeCreateServiceDTO
type MealItemDTO = dailyintakedto.MealItemDTO

func (s *MCPDailyIntakeService) GetMealCalories(params DailyIntakeCreateRequestDTO) (DailyIntakeCreateServiceDTO, error) {
	// Call the LLM client to get the meal calories for meal items
	mealItems := params.MealItems
	mealItemsDTO := make([]MealItemDTO, len(mealItems))
	for i, item := range mealItems {
		mealItemsDTO[i] = MealItemDTO(item)
	}
	// Form the prompt for the LLM client
	jsonMealItemsDTO, stringifyError := json.Marshal(mealItemsDTO)
	if stringifyError != nil {
		return DailyIntakeCreateServiceDTO{}, stringifyError
	}
	// The prompt should be a string that describes the meal items and asks for their calorie count
	userRoleInput := `Here is the list of meal items:
		` + string(jsonMealItemsDTO)
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
					"description": "the measurement of the meal item",
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
