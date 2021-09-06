package main

import (
	"net/http"

	"thewaytowire/conf"
	"thewaytowire/handler/tests"
	"thewaytowire/handler/user"
)

type Handlers struct {
	User  *user.Handler
	Tests *tests.Handler
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
	http.HandleFunc("/tests/dlv_inlining", h.Tests.DelveInlining)
}
