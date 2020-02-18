package model

type TeamMembershipCreateRequest struct {
	// User identifier
	UserID int32 `json:"UserID"`
	// Team identifier
	TeamID int32 `json:"TeamID"`
	// Role for the user inside the team (1 for leader and 2 for regular member)
	Role int32 `json:"Role"`
}
