package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {

	ech := make(chan os.Signal, 1)
	signal.Notify(ech, os.Interrupt)

	readconfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed...")
		fmt.Printf("\n file changed: [%s] \n operation: [%s] \n", in.Name, in.Op.String())
		go runthis()
	})

	go runthis()

	<-ech
	fmt.Println("main DONE...")

}

func runthis() {
	fmt.Println("running something")
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
