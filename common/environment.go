package common

import (
	"os"
)

const (
	envPhase     = "PHASE"
	defaultPhase = "local"
)

func GetEnv(envs []string, defaultVal string) string {
	for _, env := range envs {
		e := os.Getenv(env)
		if e != "" {
			return e
		}
	}

	return defaultVal
}

func IsPhaseLocal() bool {
	p := GetEnv([]string{envPhase}, defaultPhase)

	if p == defaultPhase {
		return true
	}

	return false
}
