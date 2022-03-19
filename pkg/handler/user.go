package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) info(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	user, _ := h.services.User.Info(userId)

	fmt.Println(userId)

	c.JSON(http.StatusOK, user)
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
