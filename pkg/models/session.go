package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	SessionExpiry time.Duration = 30 * time.Minute
)

// SessionToken model
// swagger:response SessionTokenResponse
type SessionToken struct {
	ID              bson.ObjectId `json:"-"`
	Username        string
	Token           string
	CreatedDateTime time.Time
	ExpiryDateTime  time.Time
	//Values          []interface{} `json:",omitempty"`
}

type Account struct {
	ID       bson.ObjectId `json:"-"`
	Username string
	Password string
}
