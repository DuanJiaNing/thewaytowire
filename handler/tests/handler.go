package tests

import (
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (u *Handler) DelveInlining(w http.ResponseWriter, r *http.Request) {
	// 编译器会对 calc 函数进行内联，从而使得 Delve 无法在函数内打断点。
	// 内联后代码: _ = 2*3
	_ = calc(2, 3)
}

func calc(a, b int) int {
	return a * b
}
