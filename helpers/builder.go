package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gov_grant_disbursement/models"
)

// HouseholOutputdBuilder : helper functions to scan returned value on struct: HouseholdOutput
func HouseholOutputdBuilder(rows *sql.Rows) []models.HouseholdOutput {
	results := []models.HouseholdOutput{}

	for rows.Next() {
		body := models.HouseholdOutput{}
		if err := rows.Scan(&body.HouseholdID, &body.HouseholdType, &body.TotalAnnualIncome); err != nil {
			fmt.Println(err)
		}
		results = append(results, body)
	}
	defer rows.Close()

	return results
}

// HouseholdOutputDetailsBuilder : helper functions to scan returned value on struct: HouseholdOutputDetails
func HouseholdOutputDetailsBuilder(rows *sql.Rows) []models.HouseholdOutputDetails {
	results := []models.HouseholdOutputDetails{}
	for rows.Next() {
		resultJSON := []uint8{}
		if err := rows.Scan(&resultJSON); err != nil {
			fmt.Println(err)
		}

		if err := json.Unmarshal(resultJSON, &results); err != nil {
			fmt.Println(err)
		}
	}
	defer rows.Close()

	return results
}
