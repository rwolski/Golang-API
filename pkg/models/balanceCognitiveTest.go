package models

import (
	"time"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type BalanceCognitiveTest struct {
	TestID               bson.ObjectId
	BalanceSessionUUID   uuid.UUID
	ChoiceSessionUUID    uuid.UUID
	ServerUpdateDateTime time.Time
}