package model

import (
	"gopkg.in/guregu/null.v4"
	"gopkg.in/guregu/null.v4/zero"
)

type Officer struct {
	Id   null.Int    `json:"id"`
	Name null.String `json:"name"`
}

type Case struct {
	Id       null.Int    `json:"id"`
	Owner    null.String `json:"owner"`
	Color    null.String `json:"color"`
	Brand    null.String `json:"model"`
	Resolved bool        `json:"resolved"`
	Officer  *Officer    `json:"officer,omitempty"`
	Moment   zero.Time   `json:"moment"`
}
