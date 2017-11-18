package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	SessionExpiry time.Duration = 30 * time.Minute
)

type TokenResponse struct {
	Token string
}

type Session struct {
	ID              bson.ObjectId `json:"-"`
	Username        string
	Token           string
	CreatedDateTime time.Time
	ExpiryDateTime  time.Time
	Values          []interface{} `json:",omitempty"`
}

type Account struct {
	ID       bson.ObjectId
	Username string
	Password string
}
