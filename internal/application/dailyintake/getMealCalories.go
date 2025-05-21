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
	// stringify the meal items
	jsonMealItemsDTO, stringifyError := json.Marshal(mealItemsDTO)
	if stringifyError != nil {
		return DailyIntakeCreateServiceDTO{}, stringifyError
	}

	// The prompt should be a string that describes the meal items and asks for their calorie count
	systemRoleInput := `You are a nutritionist. I will provide you with a list of meal items and their measurements. Please provide the calorie count for each item in the list.`
	userRoleInput := `Here is the list of meal items:
		` + string(jsonMealItemsDTO)
	logger.Logger.Infof("Prompt for LLM: %s, %s", systemRoleInput, userRoleInput)
	promptResponse, err := s.MCPClient.CreateChatCompletion(context.Background(), systemRoleInput, userRoleInput)
	if err != nil {
		logger.Logger.Errorf("Error calling LLM client: %v", err)
		return DailyIntakeCreateServiceDTO{}, err
	}
	// Parse the response from the LLM client
	var mealCaloriesResponse MealItemDTO
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
		MealItems: []MealItemDTO{mealCaloriesResponse},
	}
	logger.Logger.Infof("Meal calories response DTO: %v", mealCaloriesResponseDTO)
	return mealCaloriesResponseDTO, nil
}
