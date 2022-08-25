package main

import (
	"fast-project-golang/migration"
	"fast-project-golang/model"
	"fast-project-golang/router"
	"fast-project-golang/tools"
	"fmt"
	"os"
)

func main() {
	encrypt := tools.SHA256([]byte(""))
	fmt.Printf(encrypt)
	tools.SetConfig()
	db := model.SetupDB()
	migration.Execution()
	r := router.SetupRouter(db)
	r.Run(":" + os.Getenv("portRun"))
}
