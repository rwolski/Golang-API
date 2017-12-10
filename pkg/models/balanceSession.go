package models

import (
	"time"
)

// BalanceSession model
// swagger:response BalanceSessionResponse
type BalanceSession struct {
	SessionID            string    `json:"sessionID" bson:"_id"`
	SessionUUID          string    `json:"sessionUuid" bson:"sessionUuid"`
	SessionName          string    `json:"sessionName" bson:"sessionName"`
	SessionDate          time.Time `json:"sessionDate" bson:"sessionDate"`
	SessionType          int       `json:"sessionType" bson:"sessionType"`
	SessionNotes         string    `json:"sessionNotes" bson:"SessionNotes"`
	SessionEvent         string
	SessionInjury        string
	SubjectWeight        float32
	PlateCode            string
	SessionUserUUID      string    `json:"sessionUserUuid" bson:"sessionUserUuid"`
	SessionAdminUUID     string    `json:"sessionAdminUuid" bson:"sessionAdminUuid"`
	SessionSiteUUID      string    `json:"sessionSiteUuid" bson:"sessionSiteUuid"`
	ServerUpdateDateTime time.Time `json:"serverUpdateDateTime,omitempty" bson:"serverUpdateDateTime"`
	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}

// BalanceStandardSession std session
// swagger:response BalanceStandardSessionResponse
type BalanceStandardSession struct {
	// in: body
	BalanceSession

	BalanceTests []BalanceStandardTest
	PathTests    []BalancePathTest
	SpellTests   []BalanceSpellTest
	LosTests     []BalanceLosTest

	ChoiceTests []BalanceCognitiveTest
}
