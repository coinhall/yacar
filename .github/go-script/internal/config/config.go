package config

import (
	"log"
	"os"
)

func GetRootDir() string {
	rootDir, isPresent := os.LookupEnv("ROOT_DIR")
	if !isPresent {
		log.Fatalf("ROOT_DIR env var is not set")
	}

	return rootDir
}
