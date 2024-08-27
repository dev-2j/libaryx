// https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

package config

import (
	"os"
)

const (
	ENV_NAME_DEV = `-dev`
	ENV_NAME_PRD = ``
)

// UAT or PRD
var IsProduction = os.Getenv("GO_ENV") == "Production"

var IsEnvDEV = os.Getenv("GO_ENV") != "Production"
var IsEnvSIT = os.Getenv("GO_ENV") == "Production" && GetEnv("ENV_NAME") == ENV_NAME_DEV

var GetServiceName = func() string {
	return _Config.GetServiceName()
}

func SetServiceName(serviceName string) {
	_Config.SetServiceName(serviceName)
}

var ENV_NAME = func() string {
	if IsEnvDEV {
		return `-dev`
	}
	return GetEnv("ENV_NAME")
}()

func IsServiceModule(moduleName string) bool {
	if IsEnvDEV || GetServiceName() == `go-private` || GetServiceName() == `go-`+moduleName {
		return true
	}
	return false
}

func IsShowLogs() bool {
	return GetEnv(`show-logs`) == `1`
}
