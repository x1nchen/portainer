package model

type TeamMembership struct {
	// Membership identifier
	Id int32 `json:"Id,omitempty"`
	// User identifier
	UserID int32 `json:"UserID,omitempty"`
	// Team identifier
	TeamID int32 `json:"TeamID,omitempty"`
	// Team role (1 for team leader and 2 for team member)
	Role int32 `json:"Role,omitempty"`
}
