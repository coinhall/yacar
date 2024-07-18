package main

import (
	"bufio"
	"log"
	"os"

	"github.com/coinhall/yacar/cmd/ibcpropagator"
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

	ibcpropagator.Start(yacarFilePaths)
	sorter.Start(yacarFilePaths)
	validator.Start(yacarFilePaths, ignoreErrorSet(projRoot))
}

func ignoreErrorSet(projRoot string) map[string]struct{} {
	errorFilePath := walker.GetErrorFilePath(projRoot)
	file, err := os.Open(errorFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	set := make(map[string]struct{})
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		set[sc.Text()] = struct{}{}
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	return set
}
