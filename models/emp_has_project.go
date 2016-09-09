package models

type EmpHasProject struct {
	EmpId     *Emp     `orm:"column(emp_id);rel(fk)"`
	ProjectId *Project `orm:"column(project_id);rel(fk)"`
}
