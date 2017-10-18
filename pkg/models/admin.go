package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// Admin model
type Admin struct {
	AdminID              bson.ObjectId `json:"adminID" bson:"_id"`
	AdminUUID            uuid.UUID     `json:"adminUuid" bson:"adminUuid"`
	AdminLoginName       string        `json:"adminLoginName" bson:"adminLoginName"`
	AdminSiteUUID        uuid.UUID     `json:"adminSiteUuid" bson:"adminSiteUuid"`
	AdminStatus          int           `json:"adminStatus" bson:"adminStatus"`
	AdminActive          bool          `json:"adminActive" bson:"adminActive"`
	AdminCreationDate    time.Time     `json:"adminCreationDate" bson:"adminCreationDate"`
	AdminEmail           string        `json:"adminEmail" bson:"adminEmail"`
	AdminPassword        string        `json:"adminPassword" bson:"adminPassword"`
	AdminPasswordDate    time.Time     `json:"adminPasswordDate" bson:"adminPasswordDate"`
	LocalUpdateDateTime  time.Time     `json:"localUpdateDateTime" bson:"-"`
	ServerUpdateDateTime time.Time     `json:"serverUpdateDateTime" bson:"serverUpdateDateTime"`

	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}

// Admins is a collection of Admins
type Admins struct {
	Admins []Admin
}
