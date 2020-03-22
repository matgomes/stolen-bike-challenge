package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Officer struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

type Bike struct {
	OwnerName string `json:"owner_name" bson:"owner_name"`
	Color     string `json:"color" bson:"color"`
	Model     string `json:"model" bson:"model"`
}

type Case struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Bike    Bike          `json:"bike" bson:"bike"`
	Open    bool          `json:"open" bson:"open"`
	Officer Officer       `json:"officer" bson:"officer"`
	Date    time.Time     `json:"date" bson:"date"`
}
