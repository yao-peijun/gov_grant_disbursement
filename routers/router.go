package routers

import (
	"gov_grant_disbursement/controllers"

	"github.com/astaxie/beego"
)

func init() {
	household()
	familyMember()
	grant()
}

func household() {
	beego.Router("/api/household", &controllers.HouseholdController{}, "post:CreateHousehold")
	beego.Router("/api/household/:householdID", &controllers.HouseholdController{}, "get:GetHouseHold")
	beego.Router("/api/household/all", &controllers.HouseholdController{}, "get:GetHouseHolds")
	beego.Router("/api/household/:householdID", &controllers.HouseholdController{}, "delete:DeleteHousehold")
}

func familyMember() {
	beego.Router("/api/household/:householdID/familyMember", &controllers.FamilyMemberController{}, "post:AddFamilyMember")
	beego.Router("/api/familyMember/:familyMemberID", &controllers.FamilyMemberController{}, "delete:DeleteFamilyMember")
}

func grant() {
	beego.Router("/api/grants/?:household/?:totalIncome", &controllers.GrantController{}, "get:SearchGrant")
}
