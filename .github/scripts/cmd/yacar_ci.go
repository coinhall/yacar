package main

import (
	"log"
	"os"

	"github.com/coinhall/yacar/cmd/sorter"
	"github.com/coinhall/yacar/cmd/validator"
	"github.com/coinhall/yacar/internal/walker"
)

func main() {
	projRoot := os.Getenv("ROOT_DIR")
	if projRoot == "" {
		log.Panicln("ROOT_DIR env var not set")
	}

	yacarFilePaths := walker.GetLocalYacarFiles(projRoot)

	validator.Start(yacarFilePaths)
	sorter.Start(yacarFilePaths)
}
