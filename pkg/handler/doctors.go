package handler

import (
	"log"
	"swe/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAvailableRecords(c *gin.Context) {
	var req model.DoctorRecordReq
	if err := req.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	records, err := h.services.Doctor.GetAvailableRecords(req.DoctorID, req.Time)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(records, c)
}

func (h *Handler) createAppointment(c *gin.Context) {
	var req model.Record
	if err := req.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	log.Print(req)
	if err := h.services.Doctor.CreateRecord(&req); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}
