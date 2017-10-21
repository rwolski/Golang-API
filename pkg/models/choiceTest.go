package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// ChoiceTest model for regular tests/recall
type ChoiceTest struct {
	TestID                 bson.ObjectId `json:"testId" bson:"_id"`
	TestUUID               uuid.UUID     `json:"testUuid" bson:"testUuid"`
	TestSessionUUID        uuid.UUID     `json:"testSessionUuid" bson:"testSessionUuid"`
	TestNumber             int           `json:"testNumber" bson:"testNumber"`
	TestPattern            string        `json:"testPattern" bson:"testPattern"`
	TestCardA              string        `json:"testCardA" bson:"testCardA"`
	TestCardB              string        `json:"testCardB" bson:"testCardB"`
	TestData               []byte        `json:"testData" bson:"testData"`
	TestDifficulty         int           `json:"testDifficulty" bson:"testDifficulty"`
	TestIndex              int           `json:"testIndex" bson:"testIndex"`
	TestAccuracy           float32       `json:"testAccuracy" bson:"testAccuracy"`
	TestTotalTime          float32       `json:"testTotalTime" bson:"testTotalTime"`
	TestNumInputs          int           `json:"testNumInputs" bson:"testNumInputs"`
	TestAvgInputTime       float32       `json:"testAvgInputTime" bson:"testAvgInputTime"`
	TestNumCorrect         int           `json:"testNumCorrect" bson:"testNumCorrect"`
	TestMinInput           float32       `json:"testMinInput" bson:"testMinInput"`
	TestMaxInput           float32       `json:"testMaxInput" bson:"testMaxInput"`
	TestMedianInputTime    float32       `json:"testMedianInputTime" bson:"testMedianInputTime"`
	TestDeviationInputTime float32       `json:"testDeviationInputTime" bson:"testDeviationInputTime"`
	ServerUpdateDateTime   time.Time     `json:"serverUpdateDateTime,omitempty" bson:"serverUpdateDateTime"`
	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}
