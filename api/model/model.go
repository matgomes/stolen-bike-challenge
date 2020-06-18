package model

import (
	"time"
)

type Officer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Case struct {
	Id       int       `json:"id"`
	Owner    string    `json:"owner"`
	Color    string    `json:"color"`
	Brand    string    `json:"model"`
	Resolved bool      `json:"resolved"`
	Officer  *Officer  `json:"officer,omitempty"`
	Moment   time.Time `json:"moment"`
}
