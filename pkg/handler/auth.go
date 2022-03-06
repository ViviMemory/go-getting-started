package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"net/http"
)

func (h *Handler) authCheck(c *gin.Context) {
	var input authCheckInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, _ := h.services.Authentication.CheckAuth(input.Phone)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signUp(c *gin.Context) {
	var input model.SignUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.Authentication.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.GenerateToken(input.Name, input.Phone)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input model.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authentication.GenerateToken(input.Name, input.Phone)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

type authCheckInput struct {
	Phone string `json:"phone" binding:"required"`
}
