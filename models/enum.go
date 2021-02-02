package models

// HouseholdStruct : Household struct
type HouseholdStruct struct {
	Landed      string
	Condominium string
	HDB         string
}

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

// Household : Household type
func Household() HouseholdStruct {
	householdType := HouseholdStruct{
		Landed:      "landed",
		Condominium: "condominium",
		HDB:         "hdb",
	}
	return householdType
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
