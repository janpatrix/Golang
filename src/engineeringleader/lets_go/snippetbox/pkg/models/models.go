package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrDuplicateEmail = errors.New("models: duplicate email")

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
