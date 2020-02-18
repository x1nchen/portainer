package model

type LicenseInformation struct {
	// License key
	LicenseKey string `json:"LicenseKey,omitempty"`
	// Company associated to the license
	Company string `json:"Company,omitempty"`
	// License expiry date
	Expiration string `json:"Expiration,omitempty"`
	// Is the license valid
	Valid bool `json:"Valid,omitempty"`
}
