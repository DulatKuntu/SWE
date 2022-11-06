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
	{
		user := admin.Group("/user")
		{
			user.POST("/create", h.createUser)
			user.GET("/getAll", h.getAllUsers)
			user.GET("/get/:id", h.getUserById)
			user.POST("/update", h.updateUser)
		}
	}
	return router
}
