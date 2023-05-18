package handler

import (
	"JWT/pkg/auth"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Auth(c *gin.Context) {
	var (
		authen auth.Authentication
		resp   Response
	)
	resp.rw = c.Writer
	if authen.CheckToken(c.Request.Header.Get("Authorization")) {
		resp.SetStatusOk()
		resp.rw.Write([]byte("authentication completed!"))
	} else {
		resp.SetStatusUnauthorized()
		resp.rw.Write([]byte("invalid token!"))
	}

}

func (h *Handler) GetToken(c *gin.Context) {
	var (
		authen auth.Authentication
		resp   Response
	)
	resp.rw = c.Writer
	token, _ := authen.CreateToken("root")
	resp.rw.Header().Set("Authorization", "Bearer "+token)
	resp.SetStatusOk()
	resp.rw.Write([]byte("Token creation was successful!"))
}
