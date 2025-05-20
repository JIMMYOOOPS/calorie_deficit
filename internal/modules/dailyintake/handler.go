// internal/modules/dailyintake/handler.go
package dailyintake

import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/handler"
	"calorie_deficit/internal/pkg/logger"
	"calorie_deficit/internal/types"
	"calorie_deficit/internal/utils"
	"time"

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
func (h *Handler) CreateDailyIntake(context *gin.Context) {
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
	serviceResponse, serviceError := h.Service.CreateDailyIntake(serviceRequest)
	if serviceError != nil {
		handler.ErrorResponse(context, serviceError)
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
func (h *Handler) GetDailyIntake(context *gin.Context) {
	// Use the ID to get the daily intake from parameters
	idString := context.Param("id")
	// Validate the ID
	idUint64, parseError := utils.ParseStringToUint(idString)
	if parseError != nil {
		context.JSON(400, types.AppError(errInvalidRequest))
		return
	}
	id := uint(idUint64)
	serviceResponse, serviceError := h.Service.GetDailyIntake(id)
	if serviceError != nil {
		handler.ErrorResponse(context, serviceError)
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

// GetDailyIntakesList handles the retrieval of daily intakes for all daily intakes, the handler can take a query parameter of user_id and date
func (h *Handler) GetDailyIntakesList(context *gin.Context) {
	// Get pagination parameters
	pageString := context.DefaultQuery("page", constants.DefaultPageNumber)
	pageSizeString := context.DefaultQuery("page_size", constants.DefaultPageSize)
	page, pageSize := utils.ParsePaginationParams(pageString, pageSizeString)
	// Get the query parameters
	userIDString := context.Query("user_id")
	dateString := context.Query("date")
	mealType := context.Query("meal_type")
	// Validate and parse the user ID
	var userID *uint
	if userIDString != "" {
		userIDUint64, err := utils.ParseStringToUint(userIDString)
		if err != nil {
			context.JSON(400, types.AppError(errInvalidRequest))
			return
		}
		uid := uint(userIDUint64)
		userID = &uid
	}
	// Validate the date
	var date *time.Time
	if dateString != "" {
		parsedDate, err := utils.ParseStringToDate(dateString)
		if err != nil {
			context.JSON(400, types.AppError(errInvalidRequest))
			return
		}
		date = &parsedDate
	}

	serviceResponse, serviceError := h.Service.GetDailyIntakesList(
		DailyIntakesListServiceDTO{
			UserID:   userID,
			Date:     date,
			MealType: mealType,
			Page:     page,
			PageSize: pageSize,
		},
	)
	if serviceError != nil {
		handler.ErrorResponse(context, serviceError)
		return
	}
	items := make([]DailyIntakeCreateResponseDTO, len(serviceResponse.Items))
	for i, intake := range serviceResponse.Items {
		items[i] = DailyIntakeCreateResponseDTO{
			ID:        intake.ID,
			UserID:    intake.UserID,
			Date:      intake.Date.Format(constants.IsoFormatMilSec),
			MealType:  intake.MealType,
			MealItems: mapMealItemsToResponseDTO(intake.MealItems),
			CreatedAt: intake.CreatedAt.Format(constants.IsoFormatMilSec),
			UpdatedAt: intake.UpdatedAt.Format(constants.IsoFormatMilSec),
		}
	}
	response := types.SuccessResponse[[]DailyIntakeCreateResponseDTO]{
		Message: constants.LogMessages.General.Success,
		Data:    items,
		Meta: &types.Meta{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: serviceResponse.TotalCount,
		},
	}
	context.JSON(200, response)
}
