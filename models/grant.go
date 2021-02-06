package models

/********
read APIs
*********/

// GrantOutput : Grant output fields
type GrantOutput struct {
	StudentEncouragementBonus []HouseholdOutputDetails `json:"studentEncouragementBonus"`
	FamilyTogetherness        []HouseholdOutputDetails `json:"familyTogethernes"`
	ElderBonus                []HouseholdOutputDetails `json:"elderBonus"`
	BabySunshunGrant          []HouseholdOutputDetails `json:"babySunshunGrant"`
	YoloGstGrant              []HouseholdOutput        `json:"yoloGstGrant"`
}
