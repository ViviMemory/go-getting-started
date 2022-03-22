package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"net/http"
)

type UserWithTestOutput struct {
	Tests model.TestOutput `json:"tests"`
	User  model.UserFull   `json:"user"`
}

func (h *Handler) info(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	user, _ := h.services.User.Info(userId)

	tests, err := h.services.Test.AllTest(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var response = UserWithTestOutput{
		Tests: tests,
		User:  user,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) setRoleAdmin(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	_, err = h.services.SetRole(userId, 2)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "update role successfully",
	})
}
