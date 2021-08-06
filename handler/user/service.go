package user

import (
	"thewaytowire/conf"
	"thewaytowire/db"
)

type Service struct {
	client db.Client
	config *conf.Config
}

func NewService(client db.Client, config *conf.Config) *Service {
	return &Service{
		client: client,
		config: config,
	}
}
