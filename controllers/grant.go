package controllers

import (
	"net/http"

	"github.com/astaxie/beego"

	config "gov_grant_disbursement/configuration"
	"gov_grant_disbursement/database"
	"gov_grant_disbursement/helpers"
	"gov_grant_disbursement/models"
)

// GrantController : Grant Controller
type GrantController struct {
	beego.Controller
}

/********
read APIs
*********/

// SearchGrant : Seach grant(s) based on household size and total income
func (c *GrantController) SearchGrant() {
	household := c.Ctx.Input.Query("household")
	totalIncome := c.Ctx.Input.Query("totalIncome")
	result := models.GrantOutput{}

	// student encouragement bonus (5 < age <= 16) and annual income less than 150,000
	rows, err := database.SQLQuery(
		`SELECT 				
			json_arrayagg(
				json_object(
					'householdID', h.householdID, 
					'householdType', h.householdType,
					'familyMembers', (
						SELECT json_arrayagg(
							json_object(
								'familyMemberID', f.familyMemberID,
								'name', f.name, 
								'gender', f.gender, 
								'maritalStatus', f.maritalStatus, 
								'occupationType', f.occupationType, 
								'DOB', f.DOB, 
								'annualIncome', f.annualIncome
							)
						)
						FROM familyMember f
						WHERE f.householdID = h.householdID AND DATEDIFF(CURDATE(), f.DOB)/365<=? 
						AND DATEDIFF(CURDATE(), f.DOB)/365>?
					)
				)
			)
		FROM  household h

		WHERE

		h.householdID IN (
			SELECT f.householdID
			FROM familyMember f
			GROUP BY f.householdID
			HAVING COUNT(f.householdID) <= ? AND SUM(f.annualIncome) <= ? AND SUM(f.annualIncome) <= ?
		)

		AND

		h.householdID IN (
			SELECT f.householdID
			FROM familyMember f
			WHERE DATEDIFF(CURDATE(), f.DOB)/365<=? AND DATEDIFF(CURDATE(), f.DOB)/365>?)`,
		config.GrantSchemesAge().StudentEncouragementBonus,
		config.GrantSchemesAge().BabySunshunGrant,
		household,
		totalIncome,
		config.GrantSchemesAnnualIncome().StudentEncouragementBonus,
		config.GrantSchemesAge().StudentEncouragementBonus,
		config.GrantSchemesAge().BabySunshunGrant,
	)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}
	result.StudentEncouragementBonus = helpers.HouseholdOutputDetailsBuilder(rows)

	// family togetherness scheme (5 < age <= 18 and wife/husband stay in the same household)
	rows, err = database.SQLQuery(
		`SELECT
			json_arrayagg(
				json_object(
					'householdID', h.householdID, 
					'householdType', h.householdType,
					'familyMembers', (
						SELECT json_arrayagg(
							json_object(
								'familyMemberID', f.familyMemberID,
								'name', f.name, 
								'gender', f.gender, 
								'maritalStatus', f.maritalStatus, 
								'occupationType', f.occupationType, 
								'DOB', f.DOB, 
								'annualIncome', f.annualIncome
							)
						)
						FROM familyMember f, familyMember f2
						WHERE f.householdID = h.householdID AND f2.householdID = h.householdID AND
						(f.name = f2.spouse AND f2.name = f.spouse OR f.familyMemberID = f2.spouse AND f2.familyMemberID = f.spouse)
					)
				)
			)
		FROM  household h
			
		WHERE

		h.householdID IN (
			SELECT f.householdID
			FROM familyMember f
			GROUP BY f.householdID
			HAVING COUNT(f.householdID) <= ? AND SUM(f.annualIncome) <= ?
		)

		AND

		h.householdID IN (SELECT f.householdID
			FROM familyMember f
			INNER JOIN familyMember f2 ON
			(f.name = f2.spouse AND f2.name = f.spouse) OR
			(f.familyMemberID = f2.spouse AND f2.familyMemberID = f.spouse)
		)

		AND

		h.householdID IN (SELECT f.householdID
			FROM familyMember f
			WHERE DATEDIFF(CURDATE(), f.DOB)/365<=? AND DATEDIFF(CURDATE(), f.DOB)/365>?)`,
		household,
		totalIncome,
		config.GrantSchemesAge().FamilyTogetherness,
		config.GrantSchemesAge().BabySunshunGrant,
	)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}
	result.FamilyTogetherness = helpers.HouseholdOutputDetailsBuilder(rows)

	// elderly bonus (age >= 50)
	rows, err = database.SQLQuery(
		`SELECT
			json_arrayagg(
				json_object(
					'householdID', h.householdID, 
					'householdType', h.householdType,
					'familyMembers', (
						SELECT json_arrayagg(
							json_object(
								'familyMemberID', f.familyMemberID,
								'name', f.name, 
								'gender', f.gender, 
								'maritalStatus', f.maritalStatus, 
								'occupationType', f.occupationType, 
								'DOB', f.DOB, 
								'annualIncome', f.annualIncome
							)
						)
						FROM familyMember f
						WHERE f.householdID = h.householdID AND DATEDIFF(CURDATE(), f.DOB)/365>=?
					)
				)
			)

		FROM  household h

		WHERE
				
		h.householdID IN (
			SELECT f.householdID
			FROM familyMember f
			GROUP BY f.householdID
			HAVING COUNT(f.householdID) <= ? AND SUM(f.annualIncome) <= ?
		)

		AND

		h.householdID IN (
			SELECT f.householdID
			FROM familyMember f
			WHERE DATEDIFF(CURDATE(), f.DOB)/365>=?)`,
		config.GrantSchemesAge().ElderBonusUpper,
		household,
		totalIncome,
		config.GrantSchemesAge().ElderBonusUpper,
	)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}
	result.ElderBonus = helpers.HouseholdOutputDetailsBuilder(rows)

	// baby sunshine grant (age <= 5)
	rows, err = database.SQLQuery(
		`SELECT
			json_arrayagg(
				json_object(
					'householdID', h.householdID, 
					'householdType', h.householdType,
					'familyMembers', (
						SELECT json_arrayagg(
							json_object(
								'familyMemberID', f.familyMemberID,
								'name', f.name, 
								'gender', f.gender, 
								'maritalStatus', f.maritalStatus, 
								'occupationType', f.occupationType, 
								'DOB', f.DOB, 
								'annualIncome', f.annualIncome
							)
						)
						FROM familyMember f
						WHERE f.householdID = h.householdID AND DATEDIFF(CURDATE(), f.DOB)/365<=?
					)
				)
			)

		FROM  household h

		WHERE
				
		h.householdID IN (
			SELECT f.householdID
			FROM familyMember f
			GROUP BY f.householdID
			HAVING COUNT(f.householdID) <= ? AND SUM(f.annualIncome) <= ?
		)

		AND

		h.householdID IN (
			SELECT f.householdID
			FROM familyMember f
			WHERE DATEDIFF(CURDATE(), f.DOB)/365<=?)`,
		config.GrantSchemesAge().BabySunshunGrant,
		household,
		totalIncome,
		config.GrantSchemesAge().BabySunshunGrant,
	)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}
	result.BabySunshunGrant = helpers.HouseholdOutputDetailsBuilder(rows)

	// yolo (total annual income <= 100,000)
	rows, err = database.SQLQuery(
		`SELECT h.householdID, h.householdType, SUM(f.annualIncome) AS 'totalAnnualIncome'
		FROM familyMember f
		INNER JOIN household h ON f.householdID = h.householdID
		GROUP BY f.householdID
		HAVING COUNT(f.householdID) <= ? AND SUM(f.annualIncome) <= ? AND SUM(f.annualIncome) <= ?`,
		household,
		totalIncome,
		config.GrantSchemesAnnualIncome().Yolo,
	)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}
	result.YoloGstGrant = helpers.HouseholOutputdBuilder(rows)

	c.Data["json"] = &result
	c.ServeJSON()
}
