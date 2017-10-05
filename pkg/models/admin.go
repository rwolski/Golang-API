package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// Admin model
type Admin struct {
	AdminID           bson.ObjectId `json:"adminID" bson:"_id"`
	AdminUUID         uuid.UUID     `json:"adminUUID" bson:"adminUUID"`
	AdminLoginName    string        `json:"adminLoginName" bson:"adminLoginName"`
	AdminSiteUUID     uuid.UUID     `json:"adminSiteUUID" bson:"adminSiteUUID"`
	AdminStatus       int           `json:"adminStatus" bson:"adminStatus"`
	AdminActive       bool          `json:"adminActive" bson:"adminActive"`
	AdminCreationDate time.Time     `json:"adminCreationDate" bson:"adminCreationDate"`
	AdminEmail        string        `json:"adminEmail" bson:"adminEmail"`
	AdminPassword     string        `json:"adminPassword" bson:"adminPassword"`
	AdminPasswordDate time.Time     `json:"adminPasswordDate" bson:"adminPasswordDate"`
}