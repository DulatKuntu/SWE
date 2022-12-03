package handler

import (
	"swe/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	admin := router.Group("/admin")
	admin.POST("/login", h.loginAdmin)
	{
		user := admin.Group("/user")
		{
			user.POST("/create", h.createUser)
			user.GET("/getAll", h.getAllUsers)
			user.GET("/get/:id", h.getUserById)
			user.POST("/update", h.updateUser)
		}
		doctor := admin.Group("/doctor")
		{
			doctor.POST("/create", h.createDoctor)
			doctor.GET("/getAll", h.getAllDoctors)
			doctor.GET("/get/:id", h.getDoctorById)
			doctor.POST("/update", h.updateDoctor)
		}
		specialization := admin.Group("/specialization")
		{
			specialization.POST("/create", h.createSpecialization)
			specialization.GET("/getAll", h.getAllSpecializations)
			specialization.GET("/get/:id", h.getSpecializationById)
			specialization.POST("/update", h.updateSpecialization)
		}
	}

	doctors := router.Group("/doctors")
	{
		doctors.POST("/getAvailableRecords", h.getAvailableRecords)
		doctors.POST("/createAppointment", h.createAppointment)
	}
	return router
}
