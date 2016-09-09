package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["xxx/controllers:EmailController"] = append(beego.GlobalControllerRouter["xxx/controllers:EmailController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:EmailController"] = append(beego.GlobalControllerRouter["xxx/controllers:EmailController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:EmailController"] = append(beego.GlobalControllerRouter["xxx/controllers:EmailController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:EmailController"] = append(beego.GlobalControllerRouter["xxx/controllers:EmailController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:EmailController"] = append(beego.GlobalControllerRouter["xxx/controllers:EmailController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:EmpController"] = append(beego.GlobalControllerRouter["xxx/controllers:EmpController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:EmpController"] = append(beego.GlobalControllerRouter["xxx/controllers:EmpController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:EmpController"] = append(beego.GlobalControllerRouter["xxx/controllers:EmpController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:EmpController"] = append(beego.GlobalControllerRouter["xxx/controllers:EmpController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:EmpController"] = append(beego.GlobalControllerRouter["xxx/controllers:EmpController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:ProjectController"] = append(beego.GlobalControllerRouter["xxx/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:ProjectController"] = append(beego.GlobalControllerRouter["xxx/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:ProjectController"] = append(beego.GlobalControllerRouter["xxx/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:ProjectController"] = append(beego.GlobalControllerRouter["xxx/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["xxx/controllers:ProjectController"] = append(beego.GlobalControllerRouter["xxx/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
