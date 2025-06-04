package env

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strings"
)

type ServiceEnvironment = string

const (
	DevelopmentEnv  = "development"
	StagingEnv      = "staging"
	ProductionEnv   = "production"
	EnvironmentName = "environment"
	GoVersionName   = "go_version"
)

var (
	envName   = "SERVICE_ENV"
	goVersion string
)

func SetFromEnvFile(filepath string) error {
	if _, err := os.Stat(filepath); err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		text = strings.TrimPrefix(text, "\uFEFF")
		vars := strings.SplitN(text, "=", 2)
		if len(vars) < 2 {
			return err
		}
		if err := os.Setenv(vars[0], vars[1]); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func Init() error {
	err := SetFromEnvFile(".env")
	if err != nil && !os.IsNotExist(err) {
		log.Printf("failed to set env file: %v\n", err)
		return err
	}

	goVersion = runtime.Version()
	return nil
}

func ServiceEnv() ServiceEnvironment {
	e := os.Getenv(envName)
	if e == "" {
		e = DevelopmentEnv
	}
	return e
}

func GetVersion() string {
	return goVersion
}

func IsDevelopment() bool {
	return ServiceEnv() == DevelopmentEnv
}

func IsStaging() bool {
	return ServiceEnv() == StagingEnv
}

func IsProduction() bool {
	return ServiceEnv() == ProductionEnv
}
