package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	config "gov_grant_disbursement/configuration"
	"gov_grant_disbursement/helpers"
	"gov_grant_disbursement/models"

	"github.com/astaxie/beego"
)

// GrantController : Grant Controller
type GrantController struct {
	beego.Controller
}

// CreateHousehold : Create Household
func (c *GrantController) CreateHousehold() {
	body := models.HouseholdInput{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.CustomAbort(http.StatusBadRequest, config.ErrorMsg().General)
	}

	if !helpers.Includes(body.HouseholdType, config.HouseholdArray()) {
		c.CustomAbort(http.StatusBadRequest, config.ErrorMsg().InvalidHousehold)
	}

	// create record in db
	res, err := helpers.SQLExec(
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

// AddFamilyMember : Add family member to household
func (c *GrantController) AddFamilyMember() {
	body := models.FamilyMemberInput{}
	householdID := c.Ctx.Input.Param(":householdID")

	// if yes, check proceed to check inputs
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.CustomAbort(http.StatusBadRequest, config.ErrorMsg().General)
	}

	// check input fields
	errorMsg, err := helpers.CheckFamilyMemberInputField(body)
	if err {
		c.CustomAbort(http.StatusBadRequest, errorMsg)
	}

	// create record in db
	res, err2 := helpers.SQLExec(
		`INSERT into familyMember(householdID, name, gender, maritalStatus, 
			spouse, occupationType, DOB,annualIncome) values(?,?,?,?,?,?,?,?)`,
		householdID,
		strings.ToLower(body.Name),
		strings.ToUpper(body.Gender),
		strings.ToLower(body.MaritalStatus),
		strings.ToLower(body.Spouse),
		strings.ToLower(body.OccupationType),
		helpers.SQLDateFormat(body.DOB),
		body.AnnualIncome,
	)

	if err2 != nil {
		c.CustomAbort(http.StatusBadRequest, err2.Error())
	}

	count, err2 := res.RowsAffected()
	id, err2 := res.LastInsertId()

	// return result
	result := map[string]interface{}{
		"id":          id,
		"rowAffected": count,
	}

	if err2 != nil {
		result = map[string]interface{}{
			"error": err2.Error(),
		}
	}

	c.Data["json"] = &result
	c.ServeJSON()
}

// GetHouseHold : Get a household with family members
func (c *GrantController) GetHouseHold() {
	householdID := c.Ctx.Input.Param(":householdID")

	// query database
	results := householdDetailsBuilder(c, householdID)

	// return results
	c.Data["json"] = &results
	c.ServeJSON()
}

// GetHouseHolds : Get all households with family members
func (c *GrantController) GetHouseHolds() {
	results := householdDetailsBuilder(c, nil)
	c.Data["json"] = &results
	c.ServeJSON()
}

// SearchGrant : Seach grant(s) based on household size and total income
func (c *GrantController) SearchGrant() {
	household := c.Ctx.Input.Query("household")
	totalIncome := c.Ctx.Input.Query("totalIncome")
	result := grantBuilder(c, household, totalIncome)
	c.Data["json"] = &result
	c.ServeJSON()
}

// DeleteHousehold : Delete household and family members
func (c *GrantController) DeleteHousehold() {
	householdID := c.Ctx.Input.Param(":householdID")

	// delete from db
	res, err := helpers.SQLExec(
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

// DeleteFamilyMember : Delete family member from household
func (c *GrantController) DeleteFamilyMember() {
	familyMemeberID := c.Ctx.Input.Param(":familyMemberID")

	// delete from db
	res, err := helpers.SQLExec(
		"DELETE from familyMember where familyMemberID=?",
		familyMemeberID,
	)
	if err != nil {
		fmt.Println(err)
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
