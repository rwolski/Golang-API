package models

import (
	"time"
)

// Group model
// swagger:response GroupResponse
type Group struct {
	GroupID              string    `json:"groupId" bson:"_id"`
	GroupUUID            string    `json:"groupUuid" bson:"groupUuid"`
	GroupName            string    `json:"groupName" bson:"groupName"`
	GroupActive          bool      `json:"groupActive" bson:"groupActive"`
	LocalUpdateDateTime  time.Time `json:"localUpdateDateTime" bson:"-"`
	ServerUpdateDateTime time.Time `json:"serverUpdateDateTime" bson:"serverUpdateDateTime"`
}

// Groups is a collection of groups
// swagger:response GroupsResponse
type Groups struct {
	// in: body
	Groups []Group
}
