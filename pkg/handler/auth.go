package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) authCheck(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {

	}

	id, _ := h.services.CheckAuth(input.Phone)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Phone string `json:"phone" binding:"required"`
}
