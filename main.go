package main

import (
	"fmt"
	"go-activeMQ/active"
	"os"
	"time"

	"github.com/spf13/viper"
)

// read config
func configInit() {
	// Get config from json
	// Get config from json
	viper.SetConfigName("config/config.test")

	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("config file error: " + err.Error())
		os.Exit(1)
	}
}

func main() {
	fmt.Println("start activeMQ client at: ", time.Now().Format("2006-01-02 15:04:05"))
	configInit()
	active.Init()
}
