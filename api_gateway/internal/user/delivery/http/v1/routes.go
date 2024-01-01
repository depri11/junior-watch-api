package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *userHandlers) Routes() {
	h.group.POST("/create", h.CreateUser)
	h.group.Any("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})
}
