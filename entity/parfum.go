package entity

import "time"

type Parfum struct {
	ID          uint
	Name        string
	Brand       string
	Description string
	CreatedAt   time.Time
}
