// internal/modules/dailyintake/handler.go
package dailyintake

import (
	dailyintakeapplication "calorie_deficit/internal/application/dailyintake"
	"calorie_deficit/internal/constants"
	dailyintakedto "calorie_deficit/internal/dto/dailyintake"
	"calorie_deficit/internal/handler"
	"calorie_deficit/internal/pkg/logger"
	"calorie_deficit/internal/types"
	"calorie_deficit/internal/utils"

	"time"

	"github.com/gin-gonic/gin"
)

type (
	DailyIntakeUpdateRequestDTO = dailyintakedto.DailyIntakeUpdateRequestDTO
)

// DailyIntakeCreateRequestDTO represents the request body for creating a daily intake
func MapMealItemsToDTO(items []MealItemRequestDTO) []MealItemDTO {
	result := make([]MealItemDTO, len(items))
	for i, item := range items {
		result[i] = MealItemDTO(item) // if fields match exactly
	}
	return result
}

type Handler struct {
	Service               *Service
	MCPDailyIntakeService *dailyintakeapplication.MCPDailyIntakeService
}

// NewHandler initializes a new Handler with the provided database connection
func NewHandler(service *Service, mcpDailyIntakeService *dailyintakeapplication.MCPDailyIntakeService) *Handler {
	return &Handler{
		Service:               service,
		MCPDailyIntakeService: mcpDailyIntakeService,
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
// CreateDailyIntake godoc
// @Summary      Create a daily intake
// @Description  Create a new daily intake record
// @Tags         daily-intake
// @Accept       json
// @Produce      json
// @Param        data  body  DailyIntakeCreateRequestDTO  true  "Daily Intake Data"
// @Success      200   {object}  DailyIntakeCreateResponseDTO
// @Failure      400   {object}  types.AppError
// @Router       /daily-intake [post]
func (h *Handler) CreateDailyIntake(context *gin.Context) {
	// Bind the request body to a struct
	var createReq DailyIntakeCreateRequestDTO
	if err := context.ShouldBindJSON(&createReq); err != nil {
		logger.Logger.Error(err.Error())
		context.JSON(400, types.AppError(errInvalidRequest))
		return
	}
	// Validate the request
	mealItems := make([]dailyintakedto.MealItemRequestDTO, len(createReq.MealItems))
	for i, item := range createReq.MealItems {
		mealItems[i] = dailyintakedto.MealItemRequestDTO(item)
	}

	serviceRequest, mcpDailyIntakeServiceError := h.MCPDailyIntakeService.GetMealCalories(createReq)
	if mcpDailyIntakeServiceError != nil {
		logger.Logger.Error(mcpDailyIntakeServiceError.Error())
		context.JSON(400, types.AppError(errInvalidRequest))
		return
	}
	// Call the service layer to create the daily intake
	serviceResponse, serviceError := h.Service.CreateDailyIntake(serviceRequest)
	if serviceError != nil {
		handler.ErrorResponse(context, serviceError)
		return
	}

	handler.SuccessResponse(context, MapDailyIntakeToResponseDTO(*serviceResponse), nil)
}

// UpdateDailyIntake handles the update of an existing daily intake record
// @Summary      Update a daily intake
// @Description  Update an existing daily intake record
// @Tags         daily-intake
// @Accept       json
// @Produce      json
// @Param        id    path  int64  true  "Daily Intake ID"
// @Param        data  body  DailyIntakeUpdateRequestDTO  true  "Daily Intake Data"
// @Success      200   {object}  DailyIntakeCreateResponseDTO
// @Failure      400   {object}  types.AppError
// @Failure      404   {object}  types.AppError
// @Router       /daily-intake/{id} [patch]
func (h *Handler) UpdateDailyIntake(context *gin.Context) {
	// Use the ID to get the daily intake from parameters
	idString := context.Param("id")
	// Validate the ID
	idUint64, parseError := utils.ParseStringToUint(idString)
	if parseError != nil {
		context.JSON(400, types.AppError(errInvalidRequest))
		return
	}
	id := uint(idUint64)
	// Bind the request body to a struct
	var updateRequest DailyIntakeUpdateRequestDTO
	if err := context.ShouldBindJSON(&updateRequest); err != nil {
		logger.Logger.Error(err.Error())
		context.JSON(400, types.AppError(errInvalidRequest))
		return
	}

	// Validate the request
	serviceRequest := DailyIntakeUpdateServiceDTO(updateRequest)
	serviceResponse, serviceError := h.Service.UpdateDailyIntake(id, serviceRequest)
	if serviceError != nil {
		handler.ErrorResponse(context, serviceError)
		return
	}
	handler.SuccessResponse(context, MapDailyIntakeToResponseDTO(*serviceResponse), nil)
}

// GetDailyIntake use the ID to get the daily intake
// @Summary      Get a daily intake
// @Description  Get a daily intake by ID
// @Tags         daily-intake
// @Accept       json
// @Produce      json
// @Param        id  path  int64  true  "Daily Intake ID"
// @Success      200   {object}  DailyIntakeCreateResponseDTO
// @Failure      400   {object}  types.AppError
// @Failure      404   {object}  types.AppError
// @Router       /daily-intake/{id} [get]
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
	handler.SuccessResponse(context, MapDailyIntakeToResponseDTO(*serviceResponse), nil)
}

// GetDailyIntakesList handles the retrieval of daily intakes for all daily intakes, the handler can take a query parameter of user_id and date
// @Summary      Get daily intakes list
// @Description  Get a list of daily intakes
// @Tags         daily-intake
// @Accept       json
// @Produce      json
// @Param        page       query  int  false  "Page number"
// @Param        page_size  query  int  false  "Page size"
// @Param        user_id    query  int  false  "User ID"
// @Param        date       query  string  false  "Date in YYYY-MM-DD format"
// @Param        meal_type  query  string  false  "Meal type (breakfast, brunch, lunch, afternoonSnack, dinner, lateNightSnack)"
// @Success      200   {object}  dailyintakedto.DailyIntakeListResponseSuccessDTO
// @Failure      400   {object}  types.AppError
// @Router       /daily-intake [get]
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
		items[i] = DailyIntakeCreateResponseDTO(intake)
	}

	handler.SuccessResponse(context, items, &types.Meta{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: serviceResponse.TotalCount,
	})
}
