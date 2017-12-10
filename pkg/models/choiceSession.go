package models

import (
	"time"
)

// ChoiceSession model
type ChoiceSession struct {
	SessionID            string    `json:"sessionID" bson:"_id"`
	SessionUUID          string    `json:"sessionUuid" bson:"sessionUuid"`
	SessionName          string    `json:"sessionName" bson:"sessionName"`
	SessionDate          string    `json:"sessionDate" bson:"sessionDate"`
	SessionType          int       `json:"sessionType" bson:"sessionType"`
	SessionNotes         string    `json:"sessionNotes" bson:"SessionNotes"`
	SessionUserUUID      string    `json:"sessionUserUuid" bson:"sessionUserUuid"`
	SessionAdminUUID     string    `json:"sessionAdminUuid" bson:"sessionAdminUuid"`
	SessionSiteUUID      string    `json:"sessionSiteUuid" bson:"sessionSiteUuid"`
	ServerUpdateDateTime time.Time `json:"serverUpdateDateTime,omitempty" bson:"serverUpdateDateTime"`
	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}

// ChoiceStandardSession std session
// swagger:response ChoiceStandardSessionResponse
type ChoiceStandardSession struct {
	// in: body
	ChoiceSession

	Tests  []ChoiceTest
	Recall []ChoiceTest
}

// ChoiceMazeSession maze session
// swagger:response ChoiceStandardSessionResponse
type ChoiceMazeSession struct {
	// in: body
	ChoiceSession

	Tests []ChoiceMazeTest
}
