package models

import (
	"time"
)

// Site model
// swagger:response SiteResponse
type Site struct {
	SiteID               string    `json:"siteID" bson:"_id"`
	SiteUUID             string    `json:"siteUUID" bson:"siteUuid"`
	SiteName             string    `json:"siteName" bson:"siteName"`
	LocalUpdateDateTime  time.Time `json:"localUpdateDateTime" bson:"-"`
	ServerUpdateDateTime time.Time `json:"serverUpdateDateTime" bson:"serverUpdateDateTime"`
}
