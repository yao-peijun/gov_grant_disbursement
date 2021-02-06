package helpers

import (
	config "gov_grant_disbursement/configuration"
	"gov_grant_disbursement/models"
	"strings"
	"time"
)

// CheckFamilyMemberInputField : Check family member input fields
func CheckFamilyMemberInputField(body models.FamilyMemberInput) (string, bool) {
	errorMsg := "Ensure input fields are correct. \n"
	error := false

	if body.Name == "" {
		errorMsg += config.ErrorMsg().EmptyName + "\n"
		error = true
	}

	// gender field
	if !Includes(body.Gender, config.GenderArray()) {
		errorMsg += config.ErrorMsg().InvalidGender + "\n"
		error = true
	}

	// marital status field
	if !Includes(body.MaritalStatus, config.MaritalStatusArray()) {
		errorMsg += config.ErrorMsg().InvalidMaritalStatus + "\n"
		error = true
	}

	// empty spouse with maritalStatus = married
	if strings.ToLower(body.MaritalStatus) == models.MaritalStatus().Married && body.Spouse == "" {
		errorMsg += config.ErrorMsg().EmptySpouse + "\n"
		error = true
	}

	// occupation field
	if !Includes(body.OccupationType, config.OccupationArray()) {
		errorMsg += config.ErrorMsg().InvalidOccupation + "\n"
		error = true
	}

	if strings.ToLower(body.OccupationType) == models.OccupationType().Employed && body.AnnualIncome == 0 {
		errorMsg += config.ErrorMsg().EmptyAnnualIncome + "\n"
		error = true
	}

	// annual income field
	if body.AnnualIncome < 0 {
		errorMsg += config.ErrorMsg().InvalidAnnualIncome + "\n"
		error = true
	}

	// dob format
	_, err := time.Parse("02-01-2006", body.DOB) // go date format 02-Day, 01-Month
	if err != nil {
		errorMsg += config.ErrorMsg().InvalidDOB
		error = true
	}

	return errorMsg, error
}
