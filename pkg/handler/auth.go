package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
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

func (h *Handler) sendSms(c *gin.Context) {

	var input authCheckInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	code := RandStringBytes(6)

	h.cache.Set(input.Phone, code, 1*time.Minute)

	sendSms(input.Phone, code)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Send sms successfully",
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

func sendSms(phone string, code string) {
	accountSid := "ACaad6ab76876e7822323bbe3f91106810"
	authToken := "97214b0dd56c113e81d4524cad8545a8"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/ACaad6ab76876e7822323bbe3f91106810/Messages.json"

	// Create possible message bodies
	quotes := "Ваш код подтверждения для входа в приложение тестирования - " + code

	// Set up rand
	rand.Seed(time.Now().Unix())

	msgData := url.Values{}
	msgData.Set("To", phone)
	msgData.Set("From", "+13607039136")
	msgData.Set("Body", quotes)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}

const letterBytes = "1234567890"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
