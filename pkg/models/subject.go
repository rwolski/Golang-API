package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// Subject model
type Subject struct {
	UserID            bson.ObjectId
	UserUUID          uuid.UUID
	UserName          string
	UserGroupUUID     uuid.UUID
	UserAdditionalID  string
	UserAdditionalID2 string
	UserStatus        int
	UserSiteUUID      uuid.UUID
	UserFirstName     string
	UserLastName      string
	UserAddress       string
	UserPhoneNumber   string
	UserEmail         string
	UserActive        bool
	UserDob           time.Time
	UserGender        int
	UserCreationDate  time.Time
	UserEthnicity     string
	UserHeight        float32
}
