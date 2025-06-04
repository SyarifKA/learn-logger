package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("ini logger")
	logger.Infoln("logger info")
}

func TestLeveling(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Traceln("ini Trace")
	logger.Debugln("ini debug")
	logger.Infoln("ini Info")
	logger.Warnln("ini Warning")
	logger.Errorln("ini Error")
}

func TestFormat(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Infoln("ini format info")
	logger.Debugln("ini format debug")
	logger.Warnln("ini format warning")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, err := os.OpenFile("aplication.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger.SetOutput(file)
	logger.SetLevel(logrus.TraceLevel)

	logger.Traceln("ini trace")
	logger.Debugln("ini debug")
	logger.Infoln("ini info")
	logger.Warnln("ini warning")
	logger.Errorln("ini error")
}

func TestField(t *testing.T) {
	logger := logrus.New()

	file, err := os.OpenFile("aplication.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger.SetOutput(file)
	logger.SetLevel(logrus.TraceLevel)

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("user_id", 123).Infoln("ini log dengan field")
	fields := map[string]interface{}{
		"user_id": 1234,
		"trx_id":  18,
	}
	logger.WithFields(fields).Infoln("ini log dengan fields")
}

func TestEntry(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("user_id", 12).Infoln("this info withfield")

	entry := logrus.NewEntry(logger)
	entry.WithField("this is entry withfield", 14).Infoln("info")
	entry.Infoln("log info with entry")
}

type SimpleHook struct{}

func (s *SimpleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel}
}

func (s *SimpleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("simple", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()

	logger.AddHook(&SimpleHook{})

	logger.Infoln("ini info")
	logger.Warnln("ini warning")
	logger.Errorln("ini error")
	logger.Fatalln("ini fatal")
}
