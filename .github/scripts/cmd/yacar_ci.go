package main

import (
	"fmt"
	"log"
	"os"

	"github.com/coinhall/yacar/cmd/sorter"
	"github.com/coinhall/yacar/cmd/validator"
	"github.com/coinhall/yacar/internal/walker"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			switch err := r.(type) {
			case error:
				log.Printf("yacar ci error: %s", err.Error())
			case string:
				log.Printf("yacar ci error: %s", err)
			case fmt.Stringer:
				log.Printf("yacar ci error: %s", err.String())
			default:
				log.Printf("yacar ci error: %#v", err)
			}
		}
	}()

	projRoot, ok := os.LookupEnv("ROOT_DIR")
	if !ok {
		panic("ROOT_DIR env var not set")
	}

	yacarFilePaths := walker.GetLocalYacarFilePaths(projRoot)

	validator.Start(yacarFilePaths)
	sorter.Start(yacarFilePaths)
}
