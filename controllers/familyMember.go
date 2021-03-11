package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	config "gov_grant_disbursement/configuration"
	"gov_grant_disbursement/database"
	"gov_grant_disbursement/helpers"
	"gov_grant_disbursement/models"

	"github.com/astaxie/beego"
)

// FamilyMemberController : FamilyMember Controllers
type FamilyMemberController struct {
	beego.Controller
}

/**********
create APIs
***********/

// AddFamilyMember : Add family member to household
func (c *FamilyMemberController) AddFamilyMember() {
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
	res, err2 := database.SQLExec(
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

/**********
delete APIs
***********/

// DeleteFamilyMember : Delete family member from household
func (c *FamilyMemberController) DeleteFamilyMember() {
	familyMemeberID := c.Ctx.Input.Param(":familyMemberID")

	// delete from db
	res, err := database.SQLExec(
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
