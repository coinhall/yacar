package main

import (
	"log"
	"os"

	"github.com/coinhall/yacar/cmd/sorter"
	"github.com/coinhall/yacar/cmd/validator"
	"github.com/coinhall/yacar/internal/walker"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			os.Exit(1)
		}
	}()

	projRoot, ok := os.LookupEnv("ROOT_DIR")
	if !ok {
		panic("ROOT_DIR env var not set")
	}

	yacarFilePaths := walker.GetLocalYacarFilePaths(projRoot)

	sorter.Start(yacarFilePaths)
	validator.Start(yacarFilePaths)
}
