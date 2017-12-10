package models

// License model
// swagger:response LicenseResponse
type License struct {
	LicenseID            string `json:"licenseId" bson:"_id"`
	LicenseUUID          string `json:"licenseUuid" bson:"licenseUuid"`
	ApplicationName      string `json:"applicationName" bson:"applicationName"`
	ApplicationVariation string `json:"applicationVariation" bson:"applicationVariation"`
	ApplicationVersion   string `json:"applicationVersion" bson:"applicationVersion"`
	SerialKey            string `json:"serialKey" bson:"serialKey"`
	MachineKey           string `json:"machineKey" bson:"machineKey"`
	ActivationDate       string `json:"activationDate" bson:"activationDate"`
}
