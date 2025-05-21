// internal/handler/response.go
package handler

// Handle the error response for the API
import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/pkg/logger"
	"calorie_deficit/internal/types"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse handles the error response for the API
func ErrorResponse(context *gin.Context, err error) {
	var appError *types.AppError
	logger.Logger.Error(context.Request.Method, context.Request.URL, err)
	if errors.As(err, &appError) {
		context.JSON(appError.Code, types.ErrorResponse{
			Code:    appError.Code,
			Message: appError.Message,
		})
	} else {
		context.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    constants.ErrInternalServerError.Code,
			Message: constants.ErrInternalServerError.Message,
		})
	}
}

// SuccessResponse handles the success response for the API
func SuccessResponse(context *gin.Context, data interface{}, meta *types.Meta) {
	context.JSON(http.StatusOK, types.SuccessResponse[any]{
		Code:    http.StatusOK,
		Message: constants.LogMessages.General.Success,
		Data:    data,
		Meta:    meta,
	})
}
