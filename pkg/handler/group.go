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
	var input GroupDetailId
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	users, err := h.services.Group.DetailGroup(input.Id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) inviteUserInGroup(c *gin.Context) {
	var input InviteUserInGroup
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	idInsert, err := h.services.Group.InviteUserInGroup(input.Id, input.Phone)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var result string
	if idInsert == 0 {
		result = "error"
	} else {
		result = "success"
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": result,
	})
}

func (h *Handler) listInviteUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	groups, err := h.services.Group.ListInviteUserInGroup(userId)

	c.JSON(http.StatusOK, groups)
}

func (h *Handler) activeInviteUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input model.GroupId
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Group.ActiveInviteUserInGroup(userId, input.Id, input.IsReject)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": "success",
	})
}

type GroupDetailId struct {
	Id int `json:"group_id"`
}

type InviteUserInGroup struct {
	Id    int    `json:"group_id"`
	Phone string `json:"phone"`
}
