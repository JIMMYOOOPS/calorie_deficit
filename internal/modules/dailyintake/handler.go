// internal/modules/dailyintake/handler.go
package dailyintake

import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/pkg/logger"
	"calorie_deficit/internal/types"
	"calorie_deficit/internal/utils"

	"errors"

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

// DailyIntakeCreateRequestDTO represents the request body for creating a daily intake
func mapMealItemsToDTO(items []MealItemRequestDTO) []MealItemDTO {
	result := make([]MealItemDTO, len(items))
	for i, item := range items {
		result[i] = MealItemDTO(item) // if fields match exactly
	}
	return result
}

// DailyIntakeCreateResponseDTO represents the response body for creating a daily intake
func mapMealItemsToResponseDTO(items []MealItem) []MealItemResponseDTO {
	result := make([]MealItemResponseDTO, len(items))
	for i, item := range items {
		result[i] = MealItemResponseDTO(item) // if fields match exactly
	}
	return result
}

// CreateDailyIntakeHandler handles the creation of a new daily intake record
func (handler *Handler) CreateDailyIntakeHandler(context *gin.Context) {
	var appError *types.AppError
	// Bind the request body to a struct
	var createReq DailyIntakeCreateRequestDTO
	if err := context.ShouldBindJSON(&createReq); err != nil {
		logger.Logger.Error(err.Error())
		context.JSON(400, types.AppError(errInvalidRequest))
		return
	}
	// Validate the request
	serviceRequest := DailyIntakeCreateServiceDTO{
		UserID:    createReq.UserID,
		Date:      createReq.Date,
		MealType:  createReq.MealType,
		MealItems: mapMealItemsToDTO(createReq.MealItems),
	}
	// Call the service layer to create the daily intake
	serviceResponse, serviceError := handler.Service.CreateDailyIntake(serviceRequest)
	if serviceError != nil {
		logger.Logger.Error(serviceError.Error())
		if errors.As(serviceError, &appError) {
			context.JSON(appError.Code, types.AppError{
				Code:    appError.Code,
				Message: appError.Message,
			})
			return
		}
		context.JSON(500, types.AppError{
			Code:    500,
			Message: serviceError.Error(),
		})
		return
	}
	// Map the service response to the response DTO
	response := DailyIntakeCreateResponseDTO{
		ID:        serviceResponse.ID,
		UserID:    serviceResponse.UserID,
		Date:      serviceResponse.Date.Format(constants.IsoFormatMilSec),
		MealType:  serviceResponse.MealType,
		MealItems: mapMealItemsToResponseDTO(serviceResponse.MealItems),
		CreatedAt: serviceResponse.CreatedAt.Format(constants.IsoFormatMilSec),
		UpdatedAt: serviceResponse.UpdatedAt.Format(constants.IsoFormatMilSec),
	}
	// Return the response
	context.JSON(201, response)
}

// GetDailyIntake use the ID to get the daily intake
func (handler *Handler) GetDailyIntake(context *gin.Context) {
	var appError *types.AppError
	// Use the ID to get the daily intake from parameters
	idString := context.Param("id")
	// Validate the ID
	idUint64, parseError := utils.ParseStringToUint(idString)
	if parseError != nil {
		context.JSON(400, types.AppError(errInvalidRequest))
		return
	}
	id := uint(idUint64)
	serviceResponse, serviceError := handler.Service.GetDailyIntake(id)
	if serviceError != nil {
		logger.Logger.Error(serviceError.Error())
		if errors.As(serviceError, &appError) {
			context.JSON(appError.Code, types.AppError{
				Code:    appError.Code,
				Message: appError.Message,
			})
			return
		}
		context.JSON(500, types.AppError{
			Code:    500,
			Message: serviceError.Error(),
		})
		return
	}
	// Map the service response to the response DTO
	response := DailyIntakeCreateResponseDTO{
		ID:        serviceResponse.ID,
		UserID:    serviceResponse.UserID,
		Date:      serviceResponse.Date.Format(constants.IsoFormatMilSec),
		MealType:  serviceResponse.MealType,
		MealItems: mapMealItemsToResponseDTO(serviceResponse.MealItems),
		CreatedAt: serviceResponse.CreatedAt.Format(constants.IsoFormatMilSec),
		UpdatedAt: serviceResponse.UpdatedAt.Format(constants.IsoFormatMilSec),
	}
	context.JSON(200, response)
}
