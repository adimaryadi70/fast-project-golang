package migration

import "fast-project-golang/model"

func Execution()  {
	db := model.SetupDB()
	db.AutoMigrate(&model.Transaction{})
	db.AutoMigrate(&model.Authentification{})
	db.AutoMigrate(&model.SessionToken{})
}