package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResponse struct {
	Code     int
	Message  string
	Response interface{}
}

func Response(c *gin.Context, response *CommonResponse) {
	c.JSON(response.Code, gin.H{
		"message":  response.Message,
		"response": response.Response,
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": message,
	})
}

func Success(response any) *CommonResponse {
	return &CommonResponse{
		http.StatusOK,
		"Success",
		response,
	}
}
