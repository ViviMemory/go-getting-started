package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"net/http"
)

func (h *Handler) testCategoriesList(c *gin.Context) {
	categories, err := h.services.TestCategory.CategoriesList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *Handler) testCategoriesAdd(c *gin.Context) {
	var input model.CategoryCreated
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.AddCategoryTest(input.Title)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
