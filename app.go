package main

import (
	"net/http"

	"thewaytowire/conf"
	"thewaytowire/handler/user"
)

type Handlers struct {
	User *user.Handler
}

func NewServer(h *Handlers, config *conf.Config) *http.Server {
	ser := &http.Server{}
	ser.Addr = config.Addr
	ser.Handler = http.DefaultServeMux

	registerHandlers(h)
	return ser
}

func registerHandlers(h *Handlers) {
	http.HandleFunc("/user/avatar", h.User.Avatar)
}
