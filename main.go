package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	_ "xxx/docs"
	_ "xxx/routers"
	"xxx/models"
	"fmt"
	"xxx/controllers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/gotest?charset=utf8")
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = true

	//testO2Oreverse()

	s := controllers.BaseAdminRouter{}
	s.NestPrepare()
}

func testInsertO2O() {
	o := orm.NewOrm()
	emp := new(models.Emp)
	emp.Name = "xxx"

	profile := new(models.Profile)
	profile.Edu = "KU"
	profile.Emp = emp

	o.Insert(emp)
	o.Insert(profile)

	fmt.Println("endFunc")
}

func testInsertO2M() {
	//o := orm.NewOrm()
	//
	//emp := new(models.Emp)
	//emp.Name = "testInsertO2M";
	//
	//profile1 := new(models.Profile)
	//profile1.Edu = "Ramkhamhaeng University"
	//
	//profile2 := new(models.Profile)
	//profile2.Edu = "Chiang Mai University"
	//
	//emp.Profile = []*models.Profile{profile1, profile2}
	//
	//o.Insert(emp)
	//o.InsertMulti(1, []*models.Profile{profile1, profile2})
	//
	//fmt.Println("endFunc")
}

func testInsertM2M () {
	//o := orm.NewOrm()
	//
	//emp := new(models.Emp)
	//emp.Name = "testInsertO2M";
	//
	//project1 := new(models.Project)
	//project1.ProjectName = "Job1"
	//
	//project2 := new(models.Project)
	//project2.ProjectName = "Job2"
	//
	//emp.Project = []*models.Project{project1, project2}
	//
	//o.Insert(emp)
	//o.InsertMulti(1, []*models.Project{project1, project2})
	//
	//fmt.Println("endFunc")
}

// reverse
func testO2M() {
	o := orm.NewOrm()

	emp := models.Emp{Id:11}
	o.LoadRelated(&emp, "profile")
	o.LoadRelated(&emp, "project")

	for _, p := range emp.Project {
		fmt.Println(p.ProjectName)
	}

	for _, p := range emp.Profile {
		fmt.Println(p.Edu)
	}

	fmt.Println("endFunc")
}

func testM2O() {
	o := orm.NewOrm()

	var profile []models.Profile
	num, err := o.QueryTable(models.Profile{}).RelatedSel().All(&profile)

	for _, e := range profile {
		fmt.Println(e.Edu, e.Emp.Name)
	}

	fmt.Println(num, err)
}

func testO2Oreverse() {
	o := orm.NewOrm()

	emp := models.Emp{Id:10}
	o.LoadRelated(&emp, "email")

	fmt.Println(emp.Email.Email)

	fmt.Println("endFunc")
}

func testO2O() {
	o := orm.NewOrm()

	var email []models.Email
	num, err := o.QueryTable(models.Email{}).RelatedSel().All(&email)

	for _, e := range email {
		fmt.Println(e.Emp.Name)
	}

	fmt.Println(num, err)
}

func testM2Mreverse() {
	o := orm.NewOrm()

	emp := models.Emp{Id:11}
	o.LoadRelated(&emp, "project")

	for _, p := range emp.Project {
		fmt.Println(p.ProjectName)

	}

	fmt.Println("endFunc")
}

func testM2M()  {
	o := orm.NewOrm()

	project := models.Project{Id:2}
	o.LoadRelated(&project, "emp")

	for _, e := range project.Emp {
		fmt.Println(e.Name)

	}

	fmt.Println("endFunc")
}