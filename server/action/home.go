package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleHome is the handler for /home
func HandleHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello Gin",
	})
}
