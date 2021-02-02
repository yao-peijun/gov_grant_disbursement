package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	config "gov_grant_disbursement/configuration"
	"gov_grant_disbursement/helpers"
	"gov_grant_disbursement/models"
	"net/http"
)

// householdDetailsBuilder : Build household(s) with family members ouput format
func householdDetailsBuilder(c *GrantController, householdID interface{}) []models.HouseholdOutputDetails {
	var rows *sql.Rows
	var err error

	if householdID != nil {
		rows, err = helpers.SQLQuery(
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
							WHERE f.householdID = h.householdID
						)
					)
				)
				FROM household h
				WHERE h.householdID=?`,
			householdID,
		)
	} else {
		rows, err = helpers.SQLQuery(
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
							WHERE f.householdID = h.householdID
						)
					)
				)
				FROM household h`,
		)
	}

	return householdOutputDetailsStructBuilder(c, rows, err)
}

// grantBuilder : Build household and eligible family memebers wrt the grant schemes
func grantBuilder(c *GrantController, household string, totalIncome string) models.GrantOutput {
	result := models.GrantOutput{}

	// student encouragement bonus (5 < age <= 16) and annual income less than 150,000
	rows, err := helpers.SQLQuery(
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
	result.StudentEncouragementBonus = householdOutputDetailsStructBuilder(c, rows, err)

	// family togetherness scheme (5 < age <= 18 and wife/husband stay in the same household)
	rows, err = helpers.SQLQuery(
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
	result.FamilyTogetherness = householdOutputDetailsStructBuilder(c, rows, err)

	// elderly bonus (age >= 50)
	rows, err = helpers.SQLQuery(
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
	result.ElderBonus = householdOutputDetailsStructBuilder(c, rows, err)

	// baby sunshine grant (age <= 5)
	rows, err = helpers.SQLQuery(
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
	result.BabySunshunGrant = householdOutputDetailsStructBuilder(c, rows, err)

	// yolo (total annual income <= 100,000)
	rows, err = helpers.SQLQuery(
		`SELECT h.householdID, h.householdType, SUM(f.annualIncome) AS 'totalAnnualIncome'
		FROM familyMember f
		INNER JOIN household h ON f.householdID = h.householdID
		GROUP BY f.householdID
		HAVING COUNT(f.householdID) <= ? AND SUM(f.annualIncome) <= ? AND SUM(f.annualIncome) <= ?`,
		household,
		totalIncome,
		config.GrantSchemesAnnualIncome().Yolo,
	)
	result.YoloGstGrant = householdStructBuilder(c, rows, err)

	return result
}

// householdStructBuilder : helper functions to scan returned value on struct: HouseholdOutput
func householdStructBuilder(c *GrantController, rows *sql.Rows, err error) []models.HouseholdOutput {
	results := []models.HouseholdOutput{}

	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}

	for rows.Next() {
		body := models.HouseholdOutput{}
		err = rows.Scan(
			&body.HouseholdID,
			&body.HouseholdType,
			&body.TotalAnnualIncome,
		)
		results = append(results, body)

		if err != nil {
			fmt.Println(err)
		}
	}
	defer rows.Close()

	return results
}

// householdOutputDetailsStructBuilder : helper functions to scan returned value on struct: HouseholdOutputDetails
func householdOutputDetailsStructBuilder(c *GrantController, rows *sql.Rows, err error) []models.HouseholdOutputDetails {
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}

	results := []models.HouseholdOutputDetails{}
	for rows.Next() {
		resultJSON := []uint8{}
		if err = rows.Scan(&resultJSON); err != nil {
			fmt.Println(err)
		}

		if err := json.Unmarshal(resultJSON, &results); err != nil {
			fmt.Println(err)
		}
	}
	defer rows.Close()

	return results
}
