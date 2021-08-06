package db

type Client interface {
	Update(src interface{}) (int, error)
}
