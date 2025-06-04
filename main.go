package main

import (
	"github.com/SyarifKA/learn-logger/pkg/env"
	"github.com/SyarifKA/learn-logger/pkg/log"
)

func main() {
	// initialize environment
	err := env.Init()
	if err != nil {
		log.Fatal(err)
	}

	// initialize config log
	err = log.SetConfig(&log.Config{
		Formatter: &log.TextFormatter,
		Level:     log.TraceLevel,
		LogName:   "application.log",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("singleton debug")
	log.Info("singleton info")
	log.Warn("singleton warn")
	log.Error("singleton debug")
	log.Fatal("singleton fatal")
}
