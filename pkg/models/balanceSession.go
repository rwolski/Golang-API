package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// BalanceSession model
type BalanceSession struct {
	SessionID            bson.ObjectId `json:"sessionID" bson:"_id"`
	SessionUUID          uuid.UUID     `json:"sessionUuid" bson:"sessionUuid"`
	SessionName          string        `json:"sessionName" bson:"sessionName"`
	SessionDate          time.Time     `json:"sessionDate" bson:"sessionDate"`
	SessionType          int           `json:"sessionType" bson:"sessionType"`
	SessionNotes         string        `json:"sessionNotes" bson:"SessionNotes"`
	SessionEvent         string
	SessionInjury        string
	SubjectWeight        float32
	PlateCode            string
	SessionUserUUID      uuid.UUID `json:"sessionUserUuid" bson:"sessionUserUuid"`
	SessionAdminUUID     uuid.UUID `json:"sessionAdminUuid" bson:"sessionAdminUuid"`
	SessionSiteUUID      uuid.UUID `json:"sessionSiteUuid" bson:"sessionSiteUuid"`
	ServerUpdateDateTime time.Time `json:"serverUpdateDateTime,omitempty" bson:"serverUpdateDateTime"`
	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}

// BalanceStandardSession std session
type BalanceStandardSession struct {
	BalanceSession

	BalanceTests []BalanceStandardTest
	PathTests    []BalancePathTest
	SpellTests   []BalanceSpellTest
	LosTests     []BalanceLosTest

	ChoiceTests []BalanceChoiceKey
}
