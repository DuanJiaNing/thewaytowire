package image

import "strings"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type Metadata struct {
	Type string
}

func (s *Service) Metadata(fileName string) (*Metadata, error) {
	// 只为演示，请忽略逻辑。
	i := strings.IndexAny(fileName, ".")
	return &Metadata{Type: fileName[i:]}, nil
}
