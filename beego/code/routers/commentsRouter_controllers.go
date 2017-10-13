package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api/controllers:AgendaController"] = append(beego.GlobalControllerRouter["api/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AgendaController"] = append(beego.GlobalControllerRouter["api/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AgendaController"] = append(beego.GlobalControllerRouter["api/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AgendaController"] = append(beego.GlobalControllerRouter["api/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AgendaController"] = append(beego.GlobalControllerRouter["api/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AgendaController"] = append(beego.GlobalControllerRouter["api/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "Options",
			Router: `/`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api/controllers:AgendaController"] = append(beego.GlobalControllerRouter["api/controllers:AgendaController"],
		beego.ControllerComments{
			Method: "AgendaDeleteOptions",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"options"},
			MethodParams: param.Make(),
			Params: nil})

}
