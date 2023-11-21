package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrInvalidCredentials = errors.New("models: wrong credentials")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password []byte
	Created  time.Time
	Active   bool
}
