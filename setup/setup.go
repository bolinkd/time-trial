package setup

import "os"

const (
	envPhase     = "PHASE"
	defaultPhase = "local"
)

func EnvsString(envs []string, defaultValue string) string {
	for _, env := range envs {
		e := os.Getenv(env)
		if e != "" {
			return e
		}
	}

	return defaultValue
}

func EnvString(env string, defaultValue string) string {
	e := os.Getenv(env)
	if e != "" {
		return e
	}

	return defaultValue
}

func IsPhaseLocal() bool {
	p := EnvString(envPhase, defaultPhase)

	if p == defaultPhase || len(p) == 0 {
		return true
	}

	return false
}
