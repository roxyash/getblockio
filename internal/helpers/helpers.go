package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/roxyash/getblock/internal/types"
	"github.com/sirupsen/logrus"
)

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, types.ErrorResponse{ErrorMessage: message})
}
