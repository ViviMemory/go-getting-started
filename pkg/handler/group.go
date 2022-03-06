package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/model"
	"net/http"
)

func (h *Handler) createGroup(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.GroupInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Group.CreateGroup(input.Name, userId)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) addUser(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.GroupAddUserInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Group.AddUserInGroup(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//func (h *Handler) deleteUser(c *gin.Context) {
//	_, err := getUserId(c)
//	if err != nil {
//		return
//	}
//
//	var input model.GroupAddUserInput
//	if err := c.BindJSON(&input); err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	id, err := h.services.Group.deleteUserInGroup(input)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"id": id,
//	})
//}

type getAllGroup struct {
	Data []model.Group `json:"data"`
}

func (h *Handler) allGroup(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	groups, err := h.services.Group.GetAllGroupUser(userId)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllGroup{
		Data: groups,
	})
}

func (h *Handler) detailGroup(c *gin.Context) {

}
