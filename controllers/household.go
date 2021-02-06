package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	config "gov_grant_disbursement/configuration"
	"gov_grant_disbursement/database"
	"gov_grant_disbursement/helpers"
	"gov_grant_disbursement/models"

	"github.com/astaxie/beego"
)

// HouseholdController : Household Controller
type HouseholdController struct {
	beego.Controller
}

/**********
create APIs
***********/

// CreateHousehold : Create Household
func (c *HouseholdController) CreateHousehold() {
	body := models.HouseholdInput{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.CustomAbort(http.StatusBadRequest, config.ErrorMsg().General)
	}

	if !helpers.Includes(body.HouseholdType, config.HouseholdArray()) {
		c.CustomAbort(http.StatusBadRequest, config.ErrorMsg().InvalidHousehold)
	}

	// create record in db
	res, err := database.SQLExec(
		"INSERT into household(householdType) values(?)",
		strings.ToLower(body.HouseholdType),
	)

	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}

	count, err := res.RowsAffected()
	id, err := res.LastInsertId()

	// return result
	result := map[string]interface{}{
		"id":          id,
		"rowAffected": count,
	}

	if err != nil {
		result = map[string]interface{}{
			"error": err.Error(),
		}
	}

	c.Data["json"] = &result
	c.ServeJSON()
}

/********
read APIs
*********/

// GetHouseHold : Get a household with family members
func (c *HouseholdController) GetHouseHold() {
	householdID := c.Ctx.Input.Param(":householdID")

	// query database
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
						WHERE f.householdID = h.householdID
					)
				)
			)
			FROM household h
			WHERE h.householdID=?`,
		householdID,
	)

	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}

	results := helpers.HouseholdOutputDetailsBuilder(rows)

	// return results
	c.Data["json"] = &results
	c.ServeJSON()
}

// GetHouseHolds : Get all households with family members
func (c *HouseholdController) GetHouseHolds() {
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
						WHERE f.householdID = h.householdID
					)
				)
			)
			FROM household h`,
	)

	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}

	results := helpers.HouseholdOutputDetailsBuilder(rows)
	c.Data["json"] = &results
	c.ServeJSON()
}

/**********
delete APIs
***********/

// DeleteHousehold : Delete household and family members
func (c *HouseholdController) DeleteHousehold() {
	householdID := c.Ctx.Input.Param(":householdID")

	// delete from db
	res, err := database.SQLExec(
		"DELETE from household where householdID=?",
		householdID,
	)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, err.Error())
	}

	count, err := res.RowsAffected()

	// return result
	result := map[string]interface{}{
		"rowAffected": count,
	}

	if err != nil {
		result = map[string]interface{}{
			"error": err.Error(),
		}
	}

	c.Data["json"] = &result
	c.ServeJSON()
}
