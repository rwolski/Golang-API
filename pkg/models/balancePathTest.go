package models

import (
	"time"
)

// BalancePathTest
// swagger:response BalancePathTestResponse
type BalancePathTest struct {
	TestID               string
	TestUUID             string
	TestSessionUUID      string
	TestFormat           int
	TestType             string
	TestIndex            int
	TestResource         string
	TestOrientation      string
	TestLength           float32
	TestAvgCogD          float32
	TestAvgCogX          float32
	TestAvgCogY          float32
	TestTravelD          float32
	TestTravelX          float32
	TestTravelY          float32
	TestVelocitySpeed    float32
	TestVelocityAngle    float32
	TestSway             float64
	TestNotes            string
	TestDataPrecision    int
	TestRawData          []byte
	TestRawDataSize      int
	TestCopData          []byte
	TestCopDataSize      int
	TestDeviationX       float32
	TestDeviationY       float32
	TestMinX             float32
	TestMaxX             float32
	TestMinY             float32
	TestMaxY             float32
	TestShannonEntropyD  float32
	TestShannonEntropyX  float32
	TestShannonEntropyY  float32
	TestRenyiEntropyD    float32
	TestRenyiEntropyX    float32
	TestRenyiEntropyY    float32
	TestCollisionNum     int
	TestCollisionTime    float32
	TestThumbnailData    []byte
	TestThumbnailSize    int
	ServerUpdateDateTime time.Time
}
