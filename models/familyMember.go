package models

/*************
create inputs
**************/

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

/************
query outputs
*************/

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

/***
enum
****/

// GenderStruct : Marital status struct
type GenderStruct struct {
	F string
	M string
}

// MaritalStatusStruct : Marital status struct
type MaritalStatusStruct struct {
	Single   string
	Married  string
	Divorced string
	Windowed string
}

// OccupationStruct : Occupation struct
type OccupationStruct struct {
	Unemployed string
	Student    string
	Employed   string
}

// Gender : Gender type
func Gender() GenderStruct {
	gender := GenderStruct{
		F: "F",
		M: "M",
	}
	return gender
}

// MaritalStatus : Marital status type
func MaritalStatus() MaritalStatusStruct {
	maritalStatus := MaritalStatusStruct{
		Single:   "single",
		Married:  "married",
		Divorced: "divorced",
		Windowed: "windowed",
	}
	return maritalStatus
}

// OccupationType : Occupation type
func OccupationType() OccupationStruct {
	occupationType := OccupationStruct{
		Student:    "student",
		Employed:   "employed",
		Unemployed: "unemployed",
	}
	return occupationType
}
