package main

import (
	"log"
	"os"

	"github.com/coinhall/yacar/cmd/sorter"
	"github.com/coinhall/yacar/cmd/validator"
	"github.com/coinhall/yacar/internal/walker"
)

func main() {
	projRoot, ok := os.LookupEnv("ROOT_DIR")
	if !ok {
		log.Panicf("ROOT_DIR env var not set")
	}

	yacarFilePaths := walker.GetLocalYacarFilePaths(projRoot)

	validator.Start(yacarFilePaths)
	sorter.Start(yacarFilePaths)
}
