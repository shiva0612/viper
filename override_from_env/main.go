package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	testBindEnv()
	testAutoEnv()
}
func testAutoEnv() {
	err := godotenv.Load(".Autoenv")
	if err != nil {
		panic("error while loading environment: " + err.Error())
	}

	viper.AutomaticEnv()
	fmt.Println(viper.GetString("name"))

	// viper.SetEnvPrefix("ENV")
	// viper.AutomaticEnv()
	// fmt.Println(viper.GetString("name")) //in env: ENV_NAME

}
func testBindEnv() {
	err := godotenv.Load(".Bindenv")
	if err != nil {
		panic("error while loading environment: " + err.Error())
	}

	readconfig()

	viper.BindEnv("name")
	// viper.BindEnv("name", "NAME") // in env: NAME

	// viper.SetEnvPrefix("ENV")
	// viper.BindEnv("name", "ENV_NAME") //in env: ENV_NAME

	// viper.BindEnv("name", "TEST") //in env: TEST

	fmt.Println(viper.GetString("name"))

}

func readconfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}
