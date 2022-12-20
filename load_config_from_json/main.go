package main

import (
	"shiva/load_config_from_json/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.SugaredLogger
)

func initLogger() {
	config := zap.NewProductionConfig()
	config.Encoding = "console"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	l, _ := config.Build()
	log = l.Sugar()
}
func main() {

	initLogger()

	p := new(config.Person)

	config.LoadConfig("./person.json", p)

	log.Infof("Person = %+v", p)
	log.Infof("done...")
}
