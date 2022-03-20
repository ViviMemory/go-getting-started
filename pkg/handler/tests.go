package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"net/http"
)

func (h *Handler) CreatedTest(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.TestInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	testId, err := h.services.Test.CreatedTest(input.Title, input.CategoryId, input.AccessPrivate, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if input.AccessPrivate {
		_, err := h.services.Test.AddPrivateTestInGroup(testId, input.GroupId)
		if err != nil {
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": testId,
	})
}

func (h *Handler) AllTests(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	tests, err := h.services.Test.AllTest(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, tests)

}
