package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// Site model
type Site struct {
	SiteID               bson.ObjectId `json:"siteID" bson:"_id"`
	SiteUUID             uuid.UUID     `json:"siteUUID" bson:"siteUuid"`
	SiteName             string        `json:"siteName" bson:"siteName"`
	LocalUpdateDateTime  time.Time     `json:"localUpdateDateTime" bson:"-"`
	ServerUpdateDateTime time.Time     `json:"serverUpdateDateTime" bson:"serverUpdateDateTime"`
}
