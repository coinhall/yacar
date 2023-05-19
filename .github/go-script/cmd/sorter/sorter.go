package main

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/coinhall/yacar/internal/config"
	"github.com/coinhall/yacar/internal/enums"
	"github.com/coinhall/yacar/internal/reader"
)

func Start() {
	projRoot := config.GetRootDir()

	// Get JSONs to validate and load into memory
	chainDirsGlob := strings.Join(enums.GetAllChainNames(), ",")
	fileNamesGlob := strings.Join(enums.GetAllFileNames(), ",")
	fileGlobPattern := reader.AsGlobPattern(projRoot, chainDirsGlob, fileNamesGlob)
	filePaths, err := filepath.Glob(fileGlobPattern)
	if err != nil {
		log.Fatalf("error while getting file paths from glob pattern: %s", err)
	}

	// Sort JSONs
	sortedJSONs, err := sortYacarJSONs(filePaths)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Ensure if ordering needs to be checked, (ts vs go)
	// Ensure that fields are in order
	orderedJSONs, err := orderYacarJSONs(sortedJSONs)
	if err != nil {
		log.Fatal(err)
	}

	// Write changes to disk
	if err := writeYacarJSONs(orderedJSONs); err != nil {
		log.Fatal(err)
	}
}

func sortYacarJSONs(filePaths []string) ([]interface{}, error) {
	panic("unimplemented")
}

func orderYacarJSONs(sortedJSONs []interface{}) ([]interface{}, error) {
	panic("unimplemented")
}

func writeYacarJSONs(orderedJSONs []interface{}) error {
	panic("unimplemented")
}
