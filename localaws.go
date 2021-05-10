package main

import (
	"os"

	"github.com/subchen/go-log"
	"github.com/subchen/go-log/formatters"
	"github.com/yartat/localaws/configuration"
)

func main() {
	logger := &log.Logger{
		Level:     log.DEBUG,
		Formatter: new(formatters.TextFormatter),
		Out:       os.Stdout,
	}

	logger.Info("Start localaws")
	cfg := &configuration.LocalAwsConfiguration{}
	err := cfg.Load("./conf/localaws.yaml")
	if err != nil {
		panic(err)
	}

	logger.Info("Shutdown localaws")
}
