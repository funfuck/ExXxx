package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"xxx/models"
	"fmt"
)

type NestPreparer interface {
	NestPrepare()
}

// baseRouter implements global settings for all other routers.
type baseRouter struct {
	beego.Controller
	user    models.Emp
	isLogin bool
}
// Prepare implements Prepare method for baseRouter.
func (this *baseRouter) Prepare() {

	// page start time
	this.Data["PageStartTime"] = time.Now()

	// Setting properties.
	this.Data["AppName"] = "test"

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}

	fmt.Println("Prepare")
}
