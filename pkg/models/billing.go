package models

import (
	"time"
)

const (
	_                = iota
	SessionStarted   = iota
	SessionCompleted = iota
	SessionRejected  = iota
)

// BillingAccount model
type BillingAccount struct {
	BillingAccountID     string    `json:"billingAccountId" bson:"_id"`
	BillingAccountUUID   string    `json:"billingAccountUuid" bson:"billingAccountUuid"`
	BillingAdminUUID     string    `json:"billingAdminUuid" bson:"billingAdminUuid"`
	BillingCredits       int       `json:"billingCredits,omitempty" bson:"billingCredits"`
	BillingSiteUUID      string    `json:"billingSiteUuid" bson:"billingSiteUuid"`
	TrialAccount         bool      `json:"trialAccount" bson:"-"`
	LocalUpdateDateTime  time.Time `json:"localUpdateDateTime" bson:"-"`
	ServerUpdateDateTime time.Time `json:"serverUpdateDateTime,omitempty" bson:"serverUpdateDateTime"`
}

// BillingEvent model
type BillingEvent struct {
	BillingEventID      string    `json:"billingEventId" bson:"_id"`
	BillingEventUUID    string    `json:"billingEventUuid" bson:"billingEventUuid"`
	BillingAccountUUID  string    `json:"billingAccountUuid" bson:"billingAccountUuid"`
	SessionUUID         string    `json:"sessionUuid" bson:"sessionUuid"`
	SessionType         int       `json:"sessionType" bson:"sessionType"`
	SessionStatus       int       `json:"sessionStatus" bson:"sessionStatus"`
	LocalStartDateTime  time.Time `json:"localStartDateTime" bson:"localStartDateTime"`
	LocalEndDateTime    time.Time `json:"localEndDateTime" bson:"localEndDateTime"`
	ServerStartDateTime time.Time `json:"serverStartDateTime" bson:"serverStartDateTime"`
	ServerEndDateTime   time.Time `json:"serverEndDateTime" bson:"serverEndDateTime"`
}

// BillingAccounts is a collection of BillingAccounts
type BillingAccounts struct {
	Accounts []BillingAccount
}
