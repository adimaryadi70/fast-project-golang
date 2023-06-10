package tools

import (
	"os"
)

func SetConfig() {
	// configuration support mysql and postgres
	// DB value = mysql || postgres
	os.Setenv("DB", "postgres")
	os.Setenv("token_duration", "1")
	os.Setenv("token_secret", "adi")
	os.Setenv("portRun", "8089")
	os.Setenv("host", "localhost")
	os.Setenv("user", "postgres")
	os.Setenv("pass", "produk")
	os.Setenv("portDB", "5432")
	os.Setenv("database", "postgres")
}
