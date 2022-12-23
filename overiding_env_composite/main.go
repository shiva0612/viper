package main

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	User  string `json:"user"`
	Mongo struct {
		User string `json:"user"`
	} `json:"mongo"`
	Postgres struct {
		User string `json:"user"`
	} `json:"postgres"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	Loadconfig("./config.json")
}

func Loadconfig(path string) {
	viper.AddConfigPath(filepath.Dir(path))     //dir path
	viper.SetConfigName(filepath.Base(path))    //filename
	viper.SetConfigType(filepath.Ext(path)[1:]) //.json

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// viper.AutomaticEnv()

	config := new(Config)
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}

	fmt.Println("config: %v", viper.AllSettings())

}
