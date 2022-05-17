package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status    int         `json:"status" example:"200"`
	Data      interface{} `json:"data,omitempty" example:"{data:{customers}}"`
	Error     interface{} `json:"error,omitempty" example:"{}"`
	RequestId string      `json:"requestId" example:"3b6272b9-1ef1-45e0"`
}

func SuccessResponse(c *gin.Context, key string, body interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:    http.StatusOK,
		Data:      map[string]interface{}{key: body},
		Error:     nil,
		RequestId: c.Request.Header.Get("X-B2-Traceid"),
	})
}

func EmptySuccessResponse(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Status:    http.StatusOK,
		Data:      nil,
		Error:     nil,
		RequestId: c.Request.Header.Get("X-B2-Traceid"),
	})
}

func ErrorResponseWithCode(c *gin.Context, errorCode int, errorData *ErrorData) {
	c.JSON(errorCode, Response{
		Status:    http.StatusOK,
		Data:      nil,
		Error:     errorData,
		RequestId: c.Request.Header.Get("X-B2-Traceid"),
	})
}

func InternalServerError(c *gin.Context, message string) {
	if len(message) == 0 {
		message = "internal server error"
	}

	errorData := &ErrorData{
		Code:    INTERNAL_SERVER_ERROR,
		Message: message,
	}

	ErrorResponseWithCode(c, http.StatusInternalServerError, errorData)
}

func BadRequest(c *gin.Context, errorData interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Status:    http.StatusBadRequest,
		Data:      nil,
		Error:     errorData,
		RequestId: c.Request.Header.Get("X-B3-Traceid"),
	})
}
