package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"net/http"
)

/**
проверяем статус аутентификации
*/
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

// Создаем контейнер с временем жизни по-умолчанию равным 5 минут и удалением просроченного кеша каждые 10 минут
//cache := memorycache.New(5 * time.Minute, 10 * time.Minute)

/**
проверяем смс код в течении минуты
*/
func (h *Handler) checkSms(c *gin.Context) {

	var input authCheckSmsInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	codeFromCache, _ := h.cache.Get(input.Phone)

	fmt.Println(codeFromCache)

	if codeFromCache == input.Sms {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "check sms successfully",
		})
	}

}

/**
регистрация нового пользователя
*/
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

/**
вход по номеру телефона
*/
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

type authCheckSmsInput struct {
	Phone string `json:"phone" binding:"required"`
	Sms   string `json:"sms" binding:"required"`
}
