package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// Subject model
type Subject struct {
	SubjectID            bson.ObjectId `json:"subjectId" bson:"_id"`
	SubjectUUID          uuid.UUID     `json:"subjectUuid" bson:"subjectUuid"`
	SubjectUserName      string        `json:"subjectUserName" bson:"subjectUserName"`
	SubjectGroupUUID     uuid.UUID     `json:"subjectGroupUuid" bson:"subjectGroupUuid"`
	SubjectAdditionalID  string        `json:"subjectAdditionalId" bson:"subjectAdditionalId"`
	SubjectAdditionalID2 string        `json:"subjectAdditionalId2" bson:"subjectAdditionalId2"`
	SubjectStatus        int           `json:"subjectStatus" bson:"subjectStatus"`
	SubjectSiteUUID      uuid.UUID     `json:"subjectSiteUUID" bson:"subjectSiteUUID"`
	SubjectFirstName     string        `json:"subjectFirstName" bson:"subjectFirstName"`
	SubjectLastName      string        `json:"subjectLastName" bson:"subjectLastName"`
	SubjectAddress       string        `json:"subjectAddress" bson:"subjectAddress"`
	SubjectPhoneNumber   string        `json:"subjectPhoneNumber" bson:"subjectPhoneNumber"`
	SubjectEmail         string        `json:"subjectEmail" bson:"subjectEmail"`
	SubjectActive        bool          `json:"subjectActive" bson:"subjectActive"`
	SubjectDob           time.Time     `json:"subjectDob" bson:"subjectDob"`
	SubjectGender        int           `json:"subjectGender" bson:"subjectGender"`
	SubjectCreationDate  time.Time     `json:"subjectCreationDate" bson:"subjectCreationDate"`
	SubjectEthnicity     string        `json:"subjectEthnicity" bson:"subjectEthnicity"`
	SubjectHeight        float32       `json:"subjectHeight" bson:"subjectHeight"`
	ServerUpdateDateTime time.Time     `json:"serverUpdateDateTime" bson:"serverUpdateDateTime"`

	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}
