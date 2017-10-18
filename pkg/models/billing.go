package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// BillingAccount model
type BillingAccount struct {
	BillingAccountID     bson.ObjectId `json:"billingAccountId" bson:"_id"`
	BillingAccountUUID   uuid.UUID     `json:"billingAccountUuid" bson:"billingAccountUuid"`
	BillingAdminUUID     uuid.UUID     `json:"billingAdminUuid" bson:"billingAdminUuid"`
	BillingCredits       int           `json:"billingCredits,omitempty" bson:"billingCredits"`
	BillingSiteUUID      uuid.UUID     `json:"billingSiteUuid" bson:"billingSiteUuid"`
	LocalUpdateDateTime  time.Time     `json:"localUpdateDateTime" bson:"-"`
	ServerUpdateDateTime time.Time     `json:"serverUpdateDateTime,omitempty" bson:"serverUpdateDateTime"`
}

// BillingEvent model
type BillingEvent struct {
	BillingEventID      bson.ObjectId `json:"billingEventId" bson:"_id"`
	BillingEventUUID    uuid.UUID     `json:"billingEventUuid" bson:"billingEventUuid"`
	BillingAccountUUID  bson.ObjectId `json:"billingAccountUuid" bson:"billingAccountUuid"`
	SessionID           uuid.UUID     `json:"sessionId" bson:"sessionId"`
	SessionType         int           `json:"sessionType" bson:"sessionType"`
	SessionStatus       int           `json:"sessionStatus" bson:"sessionStatus"`
	LocalStartDateTime  time.Time     `json:"localStartDateTime" bson:"localStartDateTime"`
	LocalEndDateTime    time.Time     `json:"localEndDateTime" bson:"localEndDateTime"`
	ServerStartDateTime time.Time     `json:"serverStartDateTime" bson:"serverStartDateTime"`
	ServerEndDateTime   time.Time     `json:"serverEndDateTime" bson:"serverEndDateTime"`
}

// BillingAccounts is a collection of BillingAccounts
type BillingAccounts struct {
	Accounts []BillingAccount
}
