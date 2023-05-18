package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	rw gin.ResponseWriter
}

func (resp *Response) SetStatusOk() {
	resp.rw.WriteHeader(http.StatusOK)
	resp.rw.Write([]byte("authenticate completed."))
}

func (resp *Response) SetStatusUnauthorized() {
	resp.rw.WriteHeader(http.StatusUnauthorized)
	resp.rw.Write([]byte("invalid login or password."))
}

func (resp *Response) SetStatusBadRequest() {
	resp.rw.WriteHeader(http.StatusBadRequest)
	resp.rw.Write([]byte("please, check your request!"))
}

func (resp *Response) SetSessionCookie(cookie string) {
	resp.rw.Header().Set("Set-Cookie", cookie)
}
