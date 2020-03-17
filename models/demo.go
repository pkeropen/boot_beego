package models

import "github.com/astaxie/beego/orm"

type Demo struct {
	Id   int    `orm:"column(id);auto"`
	Name string `orm:"column(name);size(64)"`
}

func (t *Demo) TableName() string {
	return "tb_demo"
}

func init() {
	orm.RegisterModel(new(Demo))
}


