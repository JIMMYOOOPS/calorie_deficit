// internal/modules/dailyintake/handler.go
package dailyintake

import (
	"github.com/gin-gonic/gin"
)

// Handler struct to hold the database connection
type Handler struct {
	Service *Service
}

// NewHandler initializes a new Handler with the provided database connection
func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// CreateDailyIntakeHandler handles the creation of a new daily intake record
func (handler *Handler) CreateDailyIntakeHandler(context *gin.Context) {
	// Bind the request body to a struct (not shown here)
	var request DailyIntakeCreateRequestDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}
	// Validate the request
	serviceRequest := DailyIntakeCreateServiceDTO{
		UserID:    request.UserID,
		Date:      request.Date,
		MealType:  request.MealType,
		MealItems: make([]MealItemDTO, len(request.MealItems)),
	}

	// Call the service layer to create the daily intake
	serviceResponse, err := handler.Service.CreateDailyIntake(serviceRequest)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to create daily intake", "details": err.Error()})
		return
	}
	// Map the service response to the response DTO
	response := DailyIntakeCreateResponseDTO{
		UserID:    serviceResponse.UserID,
		Date:      serviceResponse.Date,
		MealType:  serviceResponse.MealType,
		MealItems: make([]MealItemResponseDTO, len(serviceResponse.MealItems)),
		CreatedAt: serviceResponse.CreatedAt,
		UpdatedAt: serviceResponse.UpdatedAt,
	}
	for i, item := range serviceResponse.MealItems {
		response.MealItems[i] = MealItemResponseDTO{
			Name:        item.Name,
			Measurement: item.Measurement,
			Quantity:    item.Quantity,
			Calorie:     item.Calorie,
		}
	}
	// Return the response
	context.JSON(201, response)
}
