package entity

import (
	"errors"
	"time"
)

var ErrAuthorNotFound = errors.New("author not found")

type Author struct {
	ID        string
	Name      string
	BookIDs   []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
