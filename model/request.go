package model

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ReqID struct
type ReqID struct {
	ID int `json:"id"`
}

func (p *ReqID) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	if p.ID == 0 {
		return errors.New("bad request | id is required")
	}

	return nil
}
