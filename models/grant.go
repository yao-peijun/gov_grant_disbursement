package models

// HouseholdInput : Household details
type HouseholdInput struct {
	HouseholdType string `json:"householdType"`
}

// FamilyMemberInput : Family memeber details fields
type FamilyMemberInput struct {
	Name           string `json:"name"`
	Gender         string `json:"gender"`
	MaritalStatus  string `json:"maritalStatus"`
	Spouse         string `json:"spouse"`
	OccupationType string `json:"occupationType"`
	AnnualIncome   int    `json:"annualIncome"`
	DOB            string `json:"DOB"`
}

// GrantOutput : Grant output fields
type GrantOutput struct {
	StudentEncouragementBonus []HouseholdOutputDetails `json:"studentEncouragementBonus"`
	FamilyTogetherness        []HouseholdOutputDetails `json:"familyTogethernes"`
	ElderBonus                []HouseholdOutputDetails `json:"elderBonus"`
	BabySunshunGrant          []HouseholdOutputDetails `json:"babySunshunGrant"`
	YoloGstGrant              []HouseholdOutput        `json:"yoloGstGrant"`
}

// HouseholdOutput : Household output fields
type HouseholdOutput struct {
	HouseholdID       int    `json:"householdID"`
	HouseholdType     string `json:"householdType"`
	TotalAnnualIncome int    `json:"totalAnnualIncome"`
}

// HouseholdOutputDetails : Household output fields
type HouseholdOutputDetails struct {
	HouseholdID   int                  `json:"householdID"`
	HouseholdType string               `json:"householdType"`
	FamilyMembers []FamilyMemberOutput `json:"familyMembers"`
}

// FamilyMemberOutput : Family memeber details fields
type FamilyMemberOutput struct {
	FamilyMemberID int    `json:"familyMemberID"`
	Name           string `json:"name"`
	Gender         string `json:"gender"`
	MaritalStatus  string `json:"maritalStatus"`
	OccupationType string `json:"occupationType"`
	AnnualIncome   int    `json:"annualIncome"`
	DOB            string `json:"DOB"`
}
