package controllers

import (
	"github.com/revel/revel"
	"github.com/revel/modules/orm/gorm/app/controllers"
)

type Users struct {
	gormcontroller.TxnController
} 

func (c Users) List() revel.Result {
	model := []Users {}
	c.Txn.Find(&model)
	return c.RenderJSON(model)
}