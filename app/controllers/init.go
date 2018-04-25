package controllers

import (
	"github.com/jinzhu/gorm"
	"../models"
	"github.com/revel/revel"
)

func initializeDB() {
	gorm.DB.AutoMigrate(&models.User{})
	var firstUser = models.User{Name: "Demo", Email: "demo@demo.com"}
	firstUser.SetNewPassword("demo")
	firstUser.Active = true
	gorm.DB.Create(&firstUser)
}

func init() {
	revel.OnAppStart(initializeDB)
}