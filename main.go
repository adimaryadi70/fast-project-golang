package main

import (
	"fast-project-golang/migration"
	"fast-project-golang/model"
	"fast-project-golang/router"
	"fast-project-golang/tools"
	"os"
)

func main() {
	tools.SetConfig()
	db := model.SetupDB()
	migration.Execution()
	r := router.SetupRouter(db)
	r.Run(":" + os.Getenv("portRun"))
}
