package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// ChoiceMazeTest model for maze tests
type ChoiceMazeTest struct {
	TestID               bson.ObjectId `json:"testId" bson:"_id"`
	TestUUID             uuid.UUID     `json:"testUuid" bson:"testUuid"`
	TestSessionUUID      uuid.UUID     `json:"testSessionUuid" bson:"testSessionUuid"`
	TestNumber           int           `json:"testNumber" bson:"testNumber"`
	TestMap              string        `json:"testMap" bson:"testMap"`
	TestScreen           string        `json:"testScreen" bson:"testScreen"`
	TestData             []byte        `json:"testData" bson:"testData"`
	TestDifficulty       int           `json:"testDifficulty" bson:"testDifficulty"`
	TestIndex            int           `json:"testIndex" bson:"testIndex"`
	TestAccuracy         float32       `json:"testAccuracy" bson:"testAccuracy"`
	TestRating           int           `json:"testRating" bson:"testRating"`
	TestTotalTime        float32       `json:"testTotalTime" bson:"testTotalTime"`
	TestCollisionTime    float32       `json:"testCollisionTime" bson:"testCollisionTime"`
	TestNumInputs        int           `json:"testNumInputs" bson:"testNumInputs"`
	TestNumCollisions    int           `json:"testNumCollisions" bson:"testNumCollisions"`
	TestMinInput         float32       `json:"testMinInput" bson:"testMinInput"`
	TestMaxInput         float32       `json:"testMaxInput" bson:"testMaxInput"`
	TestMinCollisionTime float32       `json:"testMedianInputTime" bson:"testMedianInputTime"`
	TestMaxCollisionTime float32       `json:"testDeviationInputTime" bson:"testDeviationInputTime"`
	ServerUpdateDateTime time.Time     `json:"serverUpdateDateTime,omitempty" bson:"serverUpdateDateTime"`
	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}
