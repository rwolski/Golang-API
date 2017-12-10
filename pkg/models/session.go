package models

import (
	"time"
)

const (
	SessionExpiry time.Duration = 30 * time.Minute
)

// SessionToken model
// swagger:response SessionTokenResponse
type SessionToken struct {
	ID              string `json:"-"`
	Username        string
	Token           string
	CreatedDateTime time.Time
	ExpiryDateTime  time.Time
	//Values          []interface{} `json:",omitempty"`
}

type Account struct {
	ID       string `json:"-"`
	Username string
	Password string
}
