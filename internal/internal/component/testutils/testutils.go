package testutils

import "github.com/spf13/viper"

func SetupTestingConfig() {
	viper.Set("app.env", "testing")
	viper.Set("http.port", "8080")
	viper.Set("database.dsn", "root:@tcp(127.0.0.1:3306)/example?charset=utf8mb4&parseTime=True&loc=Local")
}
