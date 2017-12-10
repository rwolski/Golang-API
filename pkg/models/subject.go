package models

import (
	"time"
)

// Subject model
// swagger:response SubjectResponse
type Subject struct {
	SubjectID            string    `json:"subjectId" bson:"_id"`
	SubjectUUID          string    `json:"subjectUuid" bson:"subjectUuid"`
	SubjectUserName      string    `json:"subjectUserName" bson:"subjectUserName"`
	SubjectGroupUUID     string    `json:"subjectGroupUuid" bson:"subjectGroupUuid"`
	SubjectAdditionalID  string    `json:"subjectAdditionalId" bson:"subjectAdditionalId"`
	SubjectAdditionalID2 string    `json:"subjectAdditionalId2" bson:"subjectAdditionalId2"`
	SubjectStatus        int       `json:"subjectStatus" bson:"subjectStatus"`
	SubjectSiteUUID      string    `json:"subjectSiteUuid" bson:"subjectSiteUuid"`
	SubjectFirstName     string    `json:"subjectFirstName" bson:"subjectFirstName"`
	SubjectLastName      string    `json:"subjectLastName" bson:"subjectLastName"`
	SubjectAddress       string    `json:"subjectAddress" bson:"subjectAddress"`
	SubjectPhoneNumber   string    `json:"subjectPhoneNumber" bson:"subjectPhoneNumber"`
	SubjectEmail         string    `json:"subjectEmail" bson:"subjectEmail"`
	SubjectActive        bool      `json:"subjectActive" bson:"subjectActive"`
	SubjectDob           time.Time `json:"subjectDob" bson:"subjectDob"`
	SubjectGender        int       `json:"subjectGender" bson:"subjectGender"`
	SubjectCreationDate  time.Time `json:"subjectCreationDate" bson:"subjectCreationDate"`
	SubjectEthnicity     string    `json:"subjectEthnicity" bson:"subjectEthnicity"`
	SubjectHeight        float32   `json:"subjectHeight" bson:"subjectHeight"`
	LocalUpdateDateTime  time.Time `json:"localUpdateDateTime" bson:"-"`
	ServerUpdateDateTime time.Time `json:"serverUpdateDateTime" bson:"serverUpdateDateTime"`
}

// Subjects is a collection of Subjects
// swagger:response SubjectsResponse
type Subjects struct {
	// in: body
	Subjects []Subject
}
