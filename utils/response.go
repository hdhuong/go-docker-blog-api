package utils

import "github.com/gin-gonic/gin"

// Response struct
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorJSON : json error response function
func ErrorJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, &Response{
		Success: false,
		Message: data.(string),
		Data:    nil,
	})
}

// SuccessJSON : json success response function
func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, &Response{
		Success: true,
		Message: data.(string),
		Data:    nil,
	})
}
