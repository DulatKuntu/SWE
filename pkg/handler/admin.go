package handler

import (
	"swe/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
	var user model.User

	if err := user.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	if err := h.services.Admin.CreateUser(&user); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(user, c)
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.Admin.GetAllUsers()

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(users, c)
}

func (h *Handler) getUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := h.services.Admin.GetUserById(id)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(user, c)
}

func (h *Handler) updateUser(c *gin.Context) {
	var user model.User

	if err := user.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	if err := h.services.Admin.UpdateUser(&user); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(user, c)
}
