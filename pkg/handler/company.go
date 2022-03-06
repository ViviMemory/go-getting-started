package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"net/http"
)

func (h *Handler) getCompany(c *gin.Context) {
	var input model.Company

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, _ := h.services.Company.GetCompany(input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}