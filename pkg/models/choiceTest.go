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
	ServerUpdateDateTime   time.Time     `json:"serverUpdateDateTime" bson:"serverUpdateDateTime"`
	//LocalUpdateDateTime  time.Time     `json:"lastLocalUpdateDateTime" bson:"lastLocalUpdateDateTime"`
}

//   <!-- List of Memory recall tests -->
//   <table name="Recall" gate="choicerecall">
//     <field name="TestId" type="int32" autoincrement="1" primary="1" desc="Test ID" gate="choicerecallid" />
//     <field name="TestUuid" type="string" desc="Flag for IsoGate sync" />
//     <field name="TestSessionId" type="int32" index="1" desc="Testing Session ID" />
//     <field name="TestSessionUuid" type="string" desc="Testing Session UUID" gate="sessionuuid" gateindex="0" />
//     <field name="TestNumber" type="int16"  index="1" desc="Test Number" gate="testnumber" gateindex="1" />
//     <field name="TestPattern" type="string" fsize="10" desc="Test Pattern" gate="testpattern" />
//     <field name="TestCardA" type="string" fsize="10" desc="Test Picture A" gate="testcarda" />
//     <field name="TestCardB" type="string" fsize="10" desc="Test Picture B" gate="testcardb" />
//     <field name="TestData" type="blob" desc="Test Data" gate="testdata" />
//     <field name="TestDifficulty" type="int32" default="1" desc="Test Difficulty Level" gate="testdifficulty" />
//     <field name="TestIndex" type="int16" desc="Baseline test index" />

//     <!-- Test calculations & statistics -->
//     <field name="TestAccuracy" type="float" desc="Test Accuracy rate" gate="testaccuracy" />
//     <field name="TestTotalTime" type="float" desc="Total Testing time" gate="testtotaltime" />
//     <field name="TestNumInputs" type="int16" desc="Number of user inputs" gate="testnuminputs" />
//     <field name="TestAvgInputTime" type="float" desc="Average time between input" gate="testavginputtime" />
//     <field name="TestNumCorrect" type="int16" default="0" desc="Number of correct inputs" />
//     <field name="TestMinInput" type="float" desc="Min input time" />
//     <field name="TestMaxInput" type="float" desc="Max input time" />
//     <field name="TestMedianInputTime" type="float" desc="Median input time" />
//     <field name="TestDeviationInputTime" type="float" desc="Std deviation of input time" />

//   </table>
