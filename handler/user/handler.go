package user

import (
	"fmt"
	"net/http"

	"thewaytowire/handler/image"
)

type Handler struct {
	s            *Service
	imageService ImageServiceProvider
}

type ImageServiceProvider interface {
	Metadata(url string) (*image.Metadata, error)
}

func NewHandler(s *Service, imageService ImageServiceProvider) *Handler {
	return &Handler{
		s:            s,
		imageService: imageService,
	}
}

func (u *Handler) Avatar(w http.ResponseWriter, r *http.Request) {
	// 只为演示，请忽略逻辑。
	file := "782c3218-b276-41a2-83c8-4ccbf1134356.jpeg"
	m, _ := u.imageService.Metadata(file)
	_, _ = w.Write([]byte(fmt.Sprintf(`file name: %s image type: %s`, file, m.Type)))
}
