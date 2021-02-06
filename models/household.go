package models

/*************
create inputs
**************/

// HouseholdInput : Household details
type HouseholdInput struct {
	HouseholdType string `json:"householdType"`
}

/************
query outputs
*************/

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

/***
enum
****/

// HouseholdStruct : Household struct
type HouseholdStruct struct {
	Landed      string
	Condominium string
	HDB         string
}

// Household : Household type
func Household() HouseholdStruct {
	householdType := HouseholdStruct{
		Landed:      "landed",
		Condominium: "condominium",
		HDB:         "hdb",
	}
	return householdType
}
