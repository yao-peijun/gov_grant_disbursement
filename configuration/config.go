package config

import (
	"gov_grant_disbursement/models"
)

// ErrorMsgStruct : Error Message struct
type ErrorMsgStruct struct {
	General              string
	DatabaseError        string
	InvalidDOB           string
	InvalidAnnualIncome  string
	InvalidHousehold     string
	InvalidOccupation    string
	InvalidGender        string
	InvalidMaritalStatus string
	EmptyName            string
	EmptySpouse          string
	EmptyAnnualIncome    string
}

// GrantSchemesAgeStruct : Grant Schemes Age struct
type GrantSchemesAgeStruct struct {
	StudentEncouragementBonus int
	FamilyTogetherness        int
	BabySunshunGrant          int
	ElderBonusUpper           int
}

// GrantSchemesAnnualIncomeStruct : GrantSchemesAnnualIncomeStruct
type GrantSchemesAnnualIncomeStruct struct {
	StudentEncouragementBonus int
	Yolo                      int
}

// HouseholdArray : Household array
func HouseholdArray() []string {
	householdType := []string{
		models.Household().Landed,
		models.Household().Condominium,
		models.Household().HDB,
	}
	return householdType
}

// GenderArray : Gender array
func GenderArray() []string {
	gender := []string{
		models.Gender().F,
		models.Gender().M,
	}
	return gender
}

// MaritalStatusArray : Marital status array
func MaritalStatusArray() []string {
	maritalStatus := []string{
		models.MaritalStatus().Single,
		models.MaritalStatus().Married,
		models.MaritalStatus().Divorced,
		models.MaritalStatus().Windowed,
	}
	return maritalStatus
}

// OccupationArray : Occupation array
func OccupationArray() []string {
	occupationType := []string{
		models.OccupationType().Employed,
		models.OccupationType().Student,
		models.OccupationType().Unemployed,
	}
	return occupationType
}

// GrantSchemesAge : Grant schemes age limit
func GrantSchemesAge() GrantSchemesAgeStruct {
	grantSchemes := GrantSchemesAgeStruct{
		StudentEncouragementBonus: 16,
		FamilyTogetherness:        18,
		BabySunshunGrant:          5,
		ElderBonusUpper:           50,
	}
	return grantSchemes
}

// GrantSchemesAnnualIncome : Grant scheme annual income limit
func GrantSchemesAnnualIncome() GrantSchemesAnnualIncomeStruct {
	grantSchemes := GrantSchemesAnnualIncomeStruct{
		StudentEncouragementBonus: 150000,
		Yolo:                      100000,
	}
	return grantSchemes
}

// ErrorMsg : list of error message
func ErrorMsg() ErrorMsgStruct {
	msg := ErrorMsgStruct{
		General:             "Error occur. Please try again.",
		DatabaseError:       "Error occur with the database. Please try again.",
		InvalidDOB:          `D.O.B should not be empty and should be in format of DD-MM-YYYY`,
		InvalidAnnualIncome: `Annual income should not be empty and should not be negative value`,
		InvalidHousehold: `Invalid household input. Input data should contain only ` +
			models.Household().Landed + `, ` + models.Household().Condominium + ` or ` + models.Household().HDB + `.`,
		InvalidOccupation: `Invalid occupation type input. Input data should contain only ` +
			models.OccupationType().Employed + `, ` + models.OccupationType().Unemployed + ` or ` + models.OccupationType().Student + `.`,
		InvalidGender: `Invalid gender input. Input data should contain only ` +
			models.Gender().F + ` or ` + models.Gender().M + `.`,
		InvalidMaritalStatus: `Invalid marital status input. Input data should contain only ` +
			models.MaritalStatus().Single + `, ` + models.MaritalStatus().Married + `, ` + models.MaritalStatus().Divorced +
			` or ` + models.MaritalStatus().Windowed + `.`,
		EmptyName:         `Name should not be empty. Please provide your name.`,
		EmptySpouse:       `Spouse's name or familyMemberID should not be empty. Please provide spouse's name or familyMemberID.`,
		EmptyAnnualIncome: `Annual income should not be empty or 0. Please provide your annual income.`,
	}
	return msg
}
