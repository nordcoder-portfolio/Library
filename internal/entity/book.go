package entity

import (
	"errors"
	"time"
)

var ErrBookNotFound = errors.New("book not found")

type Book struct {
	ID        string
	Name      string
	AuthorsID []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
