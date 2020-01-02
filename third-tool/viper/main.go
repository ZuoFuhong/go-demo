package main

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

// viper go configuration

var Conf *Config

type Config struct {
	Database  Database
	Zookeeper Zookeeper
}

type Database struct {
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Zookeeper struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

func main() {
	realPath, _ := filepath.Abs("./")
	viper.SetConfigName("db")
	viper.AddConfigPath(realPath + "/third-tool/viper")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	Conf = new(Config)
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	// 结构体序列化
	fmt.Println(Conf)

	// Key匹配
	props := viper.GetStringMap("database")
	fmt.Println(props)

	host := viper.GetString("database.url")
	fmt.Println(host)
}
