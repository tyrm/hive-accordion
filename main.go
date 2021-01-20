package main

import (
	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
	"hive-accordian/config"
	"hive-accordian/web"
	"os"
	"os/signal"
	"syscall"
)

var logger *loggo.Logger

func main() {
	conf := config.CollectConfig()

	// Init Logging
	newLogger := loggo.GetLogger("main")
	logger = &newLogger

	err := loggo.ConfigureLoggers(conf.LoggerConfig)
	if err != nil {
		logger.Errorf("Error configuring Logger: %s", err.Error())
		return
	}

	_, err = loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr))
	if err != nil {
		logger.Errorf("Error configuring Color Logger: %s", err.Error())
		return
	}

	// Init Web
	err = web.Init(conf)
	if err != nil {
		logger.Errorf("could not init web: %s", err.Error())
		return
	}

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	nch := make(chan os.Signal)
	signal.Notify(nch, syscall.SIGINT, syscall.SIGTERM)
	logger.Infof("%s", <-nch)
}
