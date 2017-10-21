package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// ChoiceSession model
type ChoiceSession struct {
	SessionID            bson.ObjectId `json:"sessionID" bson:"_id"`
	SessionUUID          uuid.UUID     `json:"sessionUuid" bson:"sessionUuid"`
	SessionName          string        `json:"sessionName" bson:"sessionName"`
	SessionDate          time.Time     `json:"sessionDate" bson:"sessionDate"`
	SessionType          int           `json:"sessionType" bson:"sessionType"`
	SessionNotes         string        `json:"sessionNotes" bson:"SessionNotes"`
	SessionUserUUID      uuid.UUID     `json:"sessionUserUuid" bson:"sessionUserUuid"`
	SessionAdminUUID     uuid.UUID     `json:"sessionAdminUuid" bson:"sessionAdminUuid"`
	SessionSiteUUID      uuid.UUID     `json:"sessionSiteUuid" bson:"sessionSiteUuid"`
	ServerUpdateDateTime time.Time     `json:"serverUpdateDateTime,omitempty" bson:"serverUpdateDateTime"`
	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}

// ChoiceStandardSession std session
type ChoiceStandardSession struct {
	ChoiceSession

	Tests  []ChoiceTest
	Recall []ChoiceTest
}

// ChoiceMazeSession maze session
type ChoiceMazeSession struct {
	ChoiceSession

	Tests []ChoiceMazeTest
}
