// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"myRestProject3/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 用http://localhost:8080/api/ 方式来访问
	beego.Router("/api", &controllers.TMCANDIDATEController{}, "get:GetAll")
	
	// 这部分urlmapping不work
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/TC_DICTIONARY1",
			beego.NSInclude(
				&controllers.TCDICTIONARY1Controller{},
			),
		),

		beego.NSNamespace("/TC_DICTIONARY2",
			beego.NSInclude(
				&controllers.TCDICTIONARY2Controller{},
			),
		),

		beego.NSNamespace("/TM_CANDIDATE",
			beego.NSInclude(
				&controllers.TMCANDIDATEController{},
			),
		),

		beego.NSNamespace("/TM_PROJECT",
			beego.NSInclude(
				&controllers.TMPROJECTController{},
			),
		),

		beego.NSNamespace("/TM_SYSTEM_USER",
			beego.NSInclude(
				&controllers.TMSYSTEMUSERController{},
			),
		),

		beego.NSNamespace("/TR_CANDIDATE_PROJECT",
			beego.NSInclude(
				&controllers.TRCANDIDATEPROJECTController{},
			),
		),

		beego.NSNamespace("/TR_CANDIDATE_USER",
			beego.NSInclude(
				&controllers.TRCANDIDATEUSERController{},
			),
		),

		beego.NSNamespace("/TR_JUDGE_BASKILL",
			beego.NSInclude(
				&controllers.TRJUDGEBASKILLController{},
			),
		),

		beego.NSNamespace("/TR_JUDGE_BATASKS",
			beego.NSInclude(
				&controllers.TRJUDGEBATASKSController{},
			),
		),

		beego.NSNamespace("/TR_JUDGE_CAPABILITY",
			beego.NSInclude(
				&controllers.TRJUDGECAPABILITYController{},
			),
		),

		beego.NSNamespace("/TR_JUDGE_MANAGE",
			beego.NSInclude(
				&controllers.TRJUDGEMANAGEController{},
			),
		),

		beego.NSNamespace("/TR_JUDGE_PROCESS",
			beego.NSInclude(
				&controllers.TRJUDGEPROCESSController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
