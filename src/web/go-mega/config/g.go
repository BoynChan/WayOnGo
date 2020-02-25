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
	// 在项目根目录和src/web/go-mega/config目录下查找配置文件
	// 设置配置文件的名字是config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath(".\\src\\web\\go-mega\\config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置失败:%s", err))
	}

}

// 获取配置文件的密钥
func GetSessionKey() string {

	return viper.GetString("session.secret-key")
}

// 获取配置文件中mysql的内容,并拼接成字符串
func GetMysqlConnectionString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true", usr, pwd, host, db, charset)
}

func GetSMTPConfig() (server string, port int, user, pwd string) {
	server = viper.GetString("mail.smtp")
	port = viper.GetInt("mail.smtp-port")
	user = viper.GetString("mail.user")
	pwd = viper.GetString("mail.password")
	return
}

func GetEmailServerUrl() string {
	return viper.GetString("mail.smtp")
}
