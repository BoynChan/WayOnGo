package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Author:Boyn
// Date:2020/2/24

func init() {
	projectName := "go-mega"
	getConfig(projectName)
}

func getConfig(projectName string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath(".\\src\\web\\go-mega\\config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置失败:%s", err))
	}

}

func GetMysqlConnectionString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true", usr, pwd, host, db, charset)
}
