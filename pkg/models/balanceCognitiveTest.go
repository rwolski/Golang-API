package models

import (
	"time"
)

type BalanceCognitiveTest struct {
	TestID               string
	BalanceSessionUUID   string
	ChoiceSessionUUID    string
	ServerUpdateDateTime time.Time
}
