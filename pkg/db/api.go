package db

type DB[T any] interface {
	Save(*T) error
	Load() (*T, error)
}
