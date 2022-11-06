package model

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type AdminLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (p *AdminLogin) ParseRequest(c *gin.Context) error {
	if err := c.BindJSON(&p); err != nil {
		return errors.New("bad request | " + err.Error())
	}

	return nil
}
