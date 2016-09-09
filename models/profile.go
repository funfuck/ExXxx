package models

import "github.com/astaxie/beego/orm"

type Profile struct {
	Id int    `orm:"column(id)"`
	Edu       string `orm:"column(edu);size(45);null"`
	Emp     *Emp   `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Profile))
}