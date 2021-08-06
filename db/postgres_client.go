package db

import "errors"

type PostgresClient struct {
}

func NewPostgresClient() (*PostgresClient, error) {
	return &PostgresClient{}, nil
}

func (c *PostgresClient) Update(src interface{}) (int, error) {
	return 0, errors.New("unimplemented")
}
