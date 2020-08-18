package model

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Officer struct {
	Id   null.Int    `json:"id"`
	Name null.String `json:"name"`
}

type Case struct {
	Id       null.Int `json:"id"`
	Resolved bool     `json:"resolved"`
	Officer  *Officer `json:"officer,omitempty"`
	CaseRequest
}

type CaseRequest struct {
	Owner  null.String `json:"owner"`
	Color  null.String `json:"color"`
	Brand  null.String `json:"brand"`
	Moment time.Time   `json:"moment"`
}
