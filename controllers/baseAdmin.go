package controllers

import "fmt"

type BaseAdminRouter struct {
	baseRouter
}

func (this *BaseAdminRouter) NestPrepare() {
	// current in admin page
	this.Data["IsAdmin"] = true
	fmt.Println("NestPrepare")
}

func (this *BaseAdminRouter) Get(){
	this.TplName = "Get.tpl"
}

func (this *BaseAdminRouter) Post(){
	this.TplName = "Post.tpl"
}