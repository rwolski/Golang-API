package models

import (
	"time"
)

// Admin model
// swagger:response AdminResponse
// in: body
type Admin struct {
	AdminID              string    `json:"adminID" bson:"_id"`
	AdminUUID            string    `json:"adminUuid" bson:"adminUuid"`
	AdminLoginName       string    `json:"adminLoginName" bson:"adminLoginName"`
	AdminSiteUUID        string    `json:"adminSiteUuid" bson:"adminSiteUuid"`
	AdminStatus          int       `json:"adminStatus" bson:"adminStatus"`
	AdminActive          bool      `json:"adminActive" bson:"adminActive"`
	AdminCreationDate    time.Time `json:"adminCreationDate" bson:"adminCreationDate"`
	AdminEmail           string    `json:"adminEmail" bson:"adminEmail"`
	AdminPassword        string    `json:"adminPassword" bson:"adminPassword"`
	AdminPasswordDate    time.Time `json:"adminPasswordDate" bson:"adminPasswordDate"`
	LocalUpdateDateTime  time.Time `json:"localUpdateDateTime" bson:"-"`
	ServerUpdateDateTime time.Time `json:"serverUpdateDateTime" bson:"serverUpdateDateTime"`
}

// Admins is a collection of Admins
// swagger:response AdminsResponse
type Admins struct {
	// in: body
	Admins []Admin
}
