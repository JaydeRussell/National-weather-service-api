package tools

import (
	"log"
	"os"
)

func MustGetEnv(name string) (env string) {
	env, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("failed to find neccessary enviroment variable: %s", name)
	}

	return
}
