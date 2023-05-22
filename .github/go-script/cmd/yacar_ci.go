package main

import (
	"log"
	"os"

	"github.com/coinhall/yacar/cmd/sorter"
	"github.com/coinhall/yacar/cmd/validator"
	"github.com/coinhall/yacar/internal/reader"
)

func main() {
	projRoot := os.Getenv("ROOT_DIR")
	if projRoot == "" {
		log.Fatalln("ROOT_DIR env var not set")
	}

	yacarFilePaths := reader.GetLocalYacarFiles(projRoot)

	validator.Start(yacarFilePaths)
	sorter.Start(yacarFilePaths)
}
