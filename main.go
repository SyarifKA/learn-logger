package main

import (
	"fmt"
	"os"
	"time"

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
	err = os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logTimestamp := time.Now().Format("2006-01-02_15-04-05")
	logFile := fmt.Sprintf("logs/%s.log", logTimestamp)

	err = log.SetConfig(&log.Config{
		Formatter: &log.TextFormatter,
		Level:     log.TraceLevel,
		LogName:   logFile,
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
