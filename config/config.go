package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type Person struct {
	Name string `mapstructure:"name"`
}

// in cls pass the *struct -> we are only returning the error if any during unmarshalling
func LoadConfig(path string, cls interface{}) error {
	viper.AddConfigPath(filepath.Dir(path))     //dir path
	viper.SetConfigName(filepath.Base(path))    //filename
	viper.SetConfigType(filepath.Ext(path)[1:]) //.json

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error while reading config file: ", path, err.Error())
		return err
	}

	// viper.GetString()
	// viper.GetInt()
	// and form struct for simple things

	// viper.AutomaticEnv() --> not working

	err = viper.Unmarshal(cls)
	if err != nil {
		log.Println("error while unmarshalling: ", err.Error())
		return err
	}

	return nil

}
