package config

import (
	"os"
	"path/filepath"
	"fmt"
	"github.com/spf13/viper"
)

// AppConfig is config object to use across application

type config struct {
	MySQLUser    string
	MySQLPass    string
	MySQLHost    string
	MySQLPort    string
	MySQLDB      string
	MySQLMaxIdle int
	MySQLMaxOpen int
	UsernameAPI       string
	PasswordAPI       string
}

var AppConfig config

func parseConfigFilePath() string {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(workPath, "config")
}
func InitializeAppConfig() {
	configPath := parseConfigFilePath()
	viper.SetConfigName("config.env")
	viper.SetConfigType("env")
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Failed to initalize appconfig.")
		return
	}
	// MySQL
	if AppConfig.MySQLUser = viper.GetString("MySQLuser"); AppConfig.MySQLUser == "" {
		panic("mysql username is missing in config")
	}
	if AppConfig.MySQLPass = viper.GetString("MySQLpass"); AppConfig.MySQLPass == "" {
		panic("mysql password is missing in config")
	}
	if AppConfig.MySQLHost = viper.GetString("MySQLhost"); AppConfig.MySQLHost == "" {
		panic("mysql host is missing in config")
	}
	if AppConfig.MySQLPort = viper.GetString("MySQLport"); AppConfig.MySQLPort == "" {
		panic("mysql port is missing in config")
	}
	if AppConfig.MySQLDB = viper.GetString("MySQLdb"); AppConfig.MySQLDB == "" {
		panic("mysql database is missing in config")
	}
	if AppConfig.MySQLMaxIdle = viper.GetInt("MySQLmaxidle"); AppConfig.MySQLMaxIdle == 0 {
		AppConfig.MySQLMaxIdle = 2
	}
	AppConfig.MySQLMaxOpen = viper.GetInt("MySQLmaxopen")

}
