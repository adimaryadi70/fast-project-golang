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
	os.Setenv("host", "43.240.224.206")
	os.Setenv("user", "hiksdi")
	os.Setenv("pass", "hiksdi2022#")
	os.Setenv("portDB", "5432")
	os.Setenv("database", "postgres")
}
