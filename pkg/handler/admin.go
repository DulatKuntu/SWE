package handler

import (
	"errors"
	"os"
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

func (h *Handler) createDoctor(c *gin.Context) {
	var doctor model.Doctor

	if err := doctor.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	if err := h.services.Admin.CreateDoctor(&doctor); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(doctor, c)
}

func (h *Handler) getAllDoctors(c *gin.Context) {
	doctors, err := h.services.Admin.GetAllDoctors()

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(doctors, c)
}

func (h *Handler) getDoctorById(c *gin.Context) {
	id := c.Param("id")

	doctor, err := h.services.Admin.GetDoctorById(id)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(doctor, c)
}

func (h *Handler) updateDoctor(c *gin.Context) {
	var doctor model.Doctor

	if err := doctor.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	if err := h.services.Admin.UpdateDoctor(&doctor); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(doctor, c)
}

func (h *Handler) loginAdmin(c *gin.Context) {
	var admin model.AdminLogin

	if err := admin.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	if admin.Login != os.Getenv("AdminLogin") || admin.Password != os.Getenv("AdminPassword") {
		defaultErrorHandler(c, errors.New("wrong login or password"))
		return
	}

	sendGeneral(admin, c)
}
