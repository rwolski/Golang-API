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
	BillingAdminID       uuid.UUID     `json:"billingAdminId" bson:"billingAdminId"`
	BillingCredits       int           `json:"billingCredits" bson:"billingCredits"`
	ServerUpdateDateTime time.Time     `json:"serverUpdateDateTime" bson:"serverUpdateDateTime"`

	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}

// BillingEvent model
type BillingEvent struct {
	BillingEventID      bson.ObjectId `json:"billingEventId" bson:"_id"`
	BillingEventUUID    uuid.UUID     `json:"billingEventUuid" bson:"billingEventUuid"`
	BillingAccountID    bson.ObjectId `json:"billingAccountId" bson:"billingAccountId"`
	SessionID           uuid.UUID     `json:"sessionId" bson:"sessionId"`
	SessionType         int           `json:"sessionType" bson:"sessionType"`
	SessionStatus       int           `json:"sessionStatus" bson:"sessionStatus"`
	LocalStartDateTime  time.Time     `json:"localStartDateTime" bson:"localStartDateTime"`
	LocalEndDateTime    time.Time     `json:"localEndDateTime" bson:"localEndDateTime"`
	ServerStartDateTime time.Time     `json:"serverStartDateTime" bson:"serverStartDateTime"`
	ServerEndDateTime   time.Time     `json:"serverEndDateTime" bson:"serverEndDateTime"`
}

// AdminBillingAccount model
// type AdminBillingAccount struct {
// 	AdminBillingID   bson.ObjectId `json:"adminBillingId" bson:"_id"`
// 	BillingAccountID bson.ObjectId `json:"billingAccountId" bson:"billingAccountId"`
// 	AdminID          bson.ObjectId `json:"adminId" bson:"adminId"`
// }
