package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	readconfig()

	// from_std_flag()
	from_viper_flag()

	fmt.Println(viper.GetString("name"))

}

func from_std_flag() {
	flag.String("name", "from flag default", "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

}

func from_viper_flag() {
	pflag.String("name", "from flag default", "help message for flagname")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
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
