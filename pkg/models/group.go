package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// Group model
type Group struct {
	GroupID              bson.ObjectId `json:"groupId" bson:"_id"`
	GroupUUID            uuid.UUID     `json:"groupUuid" bson:"groupUuid"`
	GroupName            string        `json:"groupName" bson:"groupName"`
	GroupActive          bool          `json:"groupActive" bson:"groupActive"`
	LocalUpdateDateTime  time.Time     `json:"localUpdateDateTime" bson:"-"`
	ServerUpdateDateTime time.Time     `json:"serverUpdateDateTime" bson:"serverUpdateDateTime"`

	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}

// Groups is a collection of groups
type Groups struct {
	Groups []*Group
}
