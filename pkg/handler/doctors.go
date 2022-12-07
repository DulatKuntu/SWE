package handler

import (
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
	id, err := getUserId(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	var req model.Record
	if err := req.ParseRequest(c); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	req.UserID = id
	if err := h.services.Doctor.CreateRecord(&req); err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendSuccess(c)
}

func (h *Handler) getDoctorAppointments(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	res, err := h.services.GetDoctorAppointments(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}

	sendGeneral(res, c)
}

func (h *Handler) getUserAppointments(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	res, err := h.services.GetUserAppointments(id)
	if err != nil {
		defaultErrorHandler(c, err)
		return
	}
	sendGeneral(res, c)
}
