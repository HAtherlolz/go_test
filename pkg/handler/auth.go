package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Hatherlolz/go_test"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err!= nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	
	}

    c.JSON(
		http.StatusOK,
        map[string]interface{}{
			"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	
}
