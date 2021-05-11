package main

import (
	"os"

	"flamingo.me/dingo"
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
	injector, err := dingo.NewInjector(
	)
	if err != nil {
		panic(err)
	}

	injector.Bind(log.Logger{}).ToInstance(logger)
	cfg := &configuration.LocalAwsConfiguration{}
	err = cfg.Load("./conf/localaws.yaml", injector)
	if err != nil {
		panic(err)
	}

	logger.Info("Shutdown localaws")
}
