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

	// created test in db

	testId, err := h.services.Test.CreatedTest(input.Title, 0, input.AccessPrivate, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// created all questions in db
	var questions = input.Questions
	for _, item := range questions {
		id, err := h.services.Question.AddQuestionInTest(item.Title, testId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		for _, itemAnswer := range item.Answers {
			err := h.services.Question.AddAnswerInQuestion(id, itemAnswer.Title, itemAnswer.IsRight)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}
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

type TestIdInput struct {
	Id int `json:"test_id"`
}

func (h *Handler) DetailTest(c *gin.Context) {
	var input TestIdInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.Test.DetailTest(input.Id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
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
