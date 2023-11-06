package lolapi

type Ability struct {
	Id string `json:"id"`
	AbilityLevel int `json:"abilityLevel"`
	DisplayName string `json:"displayName"`
	RawDescription string `json:"rawDescription"`
	RawDisplayName string `json:"rawDisplayName"`
}
