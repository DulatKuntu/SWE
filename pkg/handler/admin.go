package handler

import (
	"errors"
	"log"
	"os"
	"strconv"
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
	specialization := c.Request.URL.Query().Get("specialization")
	specializationID, err := strconv.Atoi(specialization)
	if err != nil {
		specializationID = 0
	}
	log.Print(specializationID)
	search := c.Request.URL.Query().Get("search")
	log.Print(search)
	doctors, err := h.services.Admin.GetAllDoctors(search, specializationID)

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

func (h *Handler) createSpecialization(c *gin.Context) {
	var specialization model.Specialization

	if err := specialization.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	if err := h.services.Admin.CreateSpecialization(&specialization); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(specialization, c)
}

func (h *Handler) getAllSpecializations(c *gin.Context) {
	specializations, err := h.services.Admin.GetAllSpecializations()

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(specializations, c)
}

func (h *Handler) getSpecializationById(c *gin.Context) {
	id := c.Param("id")

	specialization, err := h.services.Admin.GetSpecializationById(id)

	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(specialization, c)
}

func (h *Handler) updateSpecialization(c *gin.Context) {
	var specialization model.Specialization

	if err := specialization.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	if err := h.services.Admin.UpdateSpecialization(&specialization); err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(specialization, c)
}
