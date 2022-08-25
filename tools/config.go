package tools

import (
	"os"
)

func SetConfig() {
	// configuration support mysql and postgres
	// DB value = mysql || postgres
	os.Setenv("DB", "postgres")
	os.Setenv("token_duration", "2000")
	os.Setenv("token_secret", "adi")
	os.Setenv("portRun", "8089")
	os.Setenv("host", "localhost")
	os.Setenv("user", "postgres")
	os.Setenv("pass", "postgres")
	os.Setenv("portDB", "3306")
	os.Setenv("database", "postgres")
}
