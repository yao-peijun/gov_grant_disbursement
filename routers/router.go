package routers

import (
	"gov_grant_disbursement/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// create record(s)
	beego.Router("/api/household", &controllers.GrantController{}, "post:CreateHousehold")
	beego.Router("/api/household/:householdID/familyMember", &controllers.GrantController{}, "post:AddFamilyMember")

	// read record(s)
	beego.Router("/api/household/:householdID", &controllers.GrantController{}, "get:GetHouseHold")
	beego.Router("/api/household/all", &controllers.GrantController{}, "get:GetHouseHolds")
	beego.Router("/api/grants/?:household/?:totalIncome", &controllers.GrantController{}, "get:SearchGrant")

	// delete record(s)
	beego.Router("/api/household/:householdID", &controllers.GrantController{}, "delete:DeleteHousehold")
	beego.Router("/api/familyMember/:familyMemberID", &controllers.GrantController{}, "delete:DeleteFamilyMember")
}
