package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["myRestProject3/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:ObjectController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY1Controller"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY1Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY1Controller"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY1Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY1Controller"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY1Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY1Controller"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY1Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY1Controller"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY1Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY2Controller"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY2Controller"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY2Controller"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY2Controller"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY2Controller"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TCDICTIONARY2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMCANDIDATEController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMCANDIDATEController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMCANDIDATEController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMCANDIDATEController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMCANDIDATEController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMCANDIDATEController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMCANDIDATEController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMCANDIDATEController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMCANDIDATEController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMCANDIDATEController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMPROJECTController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMPROJECTController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMPROJECTController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMPROJECTController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMPROJECTController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMPROJECTController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMPROJECTController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMPROJECTController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMPROJECTController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMPROJECTController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMSYSTEMUSERController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMSYSTEMUSERController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMSYSTEMUSERController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMSYSTEMUSERController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMSYSTEMUSERController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMSYSTEMUSERController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMSYSTEMUSERController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMSYSTEMUSERController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TMSYSTEMUSERController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TMSYSTEMUSERController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEPROJECTController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEPROJECTController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEPROJECTController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEPROJECTController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEPROJECTController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEPROJECTController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEPROJECTController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEPROJECTController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEPROJECTController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEPROJECTController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEUSERController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEUSERController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEUSERController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEUSERController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEUSERController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEUSERController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEUSERController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEUSERController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEUSERController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRCANDIDATEUSERController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBASKILLController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBASKILLController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBASKILLController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBASKILLController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBASKILLController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBASKILLController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBASKILLController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBASKILLController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBASKILLController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBASKILLController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBATASKSController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBATASKSController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBATASKSController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBATASKSController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBATASKSController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBATASKSController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBATASKSController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBATASKSController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBATASKSController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEBATASKSController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGECAPABILITYController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGECAPABILITYController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGECAPABILITYController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGECAPABILITYController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGECAPABILITYController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGECAPABILITYController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGECAPABILITYController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGECAPABILITYController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGECAPABILITYController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGECAPABILITYController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEMANAGEController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEMANAGEController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEMANAGEController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEMANAGEController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEMANAGEController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEMANAGEController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEMANAGEController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEMANAGEController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEMANAGEController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEMANAGEController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEPROCESSController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEPROCESSController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEPROCESSController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEPROCESSController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEPROCESSController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEPROCESSController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEPROCESSController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEPROCESSController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEPROCESSController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:TRJUDGEPROCESSController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:UserController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:UserController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:UserController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:UserController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myRestProject3/controllers:UserController"] = append(beego.GlobalControllerRouter["myRestProject3/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
