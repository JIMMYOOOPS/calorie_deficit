// internal/modules/dailyintake/handler.go
package dailyintake

import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/pkg/logger"
	"calorie_deficit/internal/types"

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

var (
	// ErrInvalidRequest is an error for invalid request
	errInvalidRequest = types.ErrorResponse{
		Code:    400,
		Message: constants.LogMessages.General.InvalidRequest,
	}
)

// CreateDailyIntakeHandler handles the creation of a new daily intake record
func (handler *Handler) CreateDailyIntakeHandler(context *gin.Context) {
	// Bind the request body to a struct
	var createReq DailyIntakeCreateRequestDTO
	if err := context.ShouldBindJSON(&createReq); err != nil {
		logger.Logger.Error(err.Error())
		context.JSON(400, errInvalidRequest)
		return
	}
	// Validate the request
	serviceRequest := DailyIntakeCreateServiceDTO{
		UserID:    createReq.UserID,
		Date:      createReq.Date,
		MealType:  createReq.MealType,
		MealItems: make([]MealItemDTO, len(createReq.MealItems)),
	}
	for i, item := range createReq.MealItems {
		serviceRequest.MealItems[i] = MealItemDTO{
			Name:        item.Name,
			Measurement: item.Measurement,
			Quantity:    item.Quantity,
			Calorie:     item.Calorie,
		}
	}
	// Call the service layer to create the daily intake
	serviceResponse, serviceError := handler.Service.CreateDailyIntake(serviceRequest)
	if serviceError != nil {
		logger.Logger.Error(serviceError.Error())
		context.JSON(500, types.ErrorResponse{
			Code:    500,
			Message: serviceError.Error(),
		})
		return
	}
	// Map the service response to the response DTO
	response := DailyIntakeCreateResponseDTO{
		UserID:    serviceResponse.UserID,
		Date:      serviceResponse.Date.Format("2006-01-02T15:04:05.000Z"),
		MealType:  serviceResponse.MealType,
		MealItems: make([]MealItemResponseDTO, len(serviceResponse.MealItems)),
		CreatedAt: serviceResponse.CreatedAt.Format("2006-01-02T15:04:05.000Z"),
		UpdatedAt: serviceResponse.UpdatedAt.Format("2006-01-02T15:04:05.000Z"),
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
