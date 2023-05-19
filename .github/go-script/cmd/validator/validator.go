package validator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/coinhall/yacar/internal/config"
	"github.com/coinhall/yacar/internal/enums"
	"github.com/coinhall/yacar/internal/reader"
	"github.com/coinhall/yacar_util"
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

	// Validate JSONs
	if err := validateYacarJSONs(filePaths); err != nil {
		log.Fatal(err)
	}

	log.Printf("Validated the following files: %s", strings.Join(filePaths, "\n  "))
}

func validateYacarJSONs(filePaths []string) error {
	var wg sync.WaitGroup
	for _, filePath := range filePaths {
		wg.Add(1)
		go func(filePath string) error {
			defer wg.Done()
			if err := validateYacarJSON(filePath); err != nil {
				return err
			}
			return nil
		}(filePath)
	}
	wg.Wait()

	return nil
}

func validateYacarJSON(filePath string) error {
	// Load file into memory
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error while opening file: %s", err)
	}
	defer file.Close()

	// Validate JSON
	if err := validation_handler(filePath, file); err != nil {
		return fmt.Errorf("error while validating JSON: %s", err)
	}

	return nil
}

// TODO: Use validator pkg or shift the respective logic into yacar_util pkg
func validation_handler(filePath string, file *os.File) error {
	switch {
	case strings.Contains(filePath, enums.Account.Name()):
		return validateAccountJSON(file)
	case strings.Contains(filePath, enums.Asset.Name()):
		return validateAssetJSON(file)
	case strings.Contains(filePath, enums.Binary.Name()):
		return validateBinaryJSON(file)
	case strings.Contains(filePath, enums.Contract.Name()):
		return validateContractJSON(file)
	case strings.Contains(filePath, enums.Entity.Name()):
		return validateEntityJSON(file)
	case strings.Contains(filePath, enums.Pool.Name()):
		return validatePoolJSON(file)
	default:
		return fmt.Errorf("unknown file name: %s", filePath)
	}
}

func validateAccountJSON(file *os.File) error {
	var accounts []yacar_util.Account

	if err := json.NewDecoder(file).Decode(&accounts); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	idCount := make(map[string]int)
	for _, account := range accounts {
		if len(account.Entity) <= 1 || len(account.Id) <= 1 {
			return fmt.Errorf("error while validating account JSON: %v", account)
		}

		idCount[account.Id]++
		if idCount[account.Id] > 1 {
			return fmt.Errorf("duplicate account ID: %s", account.Id)
		}
	}

	return nil
}

func validateAssetJSON(file *os.File) error {
	var assets []yacar_util.Asset

	if err := json.NewDecoder(file).Decode(&assets); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	idCount := make(map[string]int)
	for _, asset := range assets {
		if !asset.IsMinimallyPopulated() {
			return fmt.Errorf("error while validating asset JSON: %v", asset)
		}

		idCount[asset.Id]++
		if idCount[asset.Id] > 1 {
			return fmt.Errorf("duplicate asset ID: %s", asset.Id)
		}
	}

	return nil
}

func validateBinaryJSON(file *os.File) error {
	var binaries []yacar_util.Binary

	if err := json.NewDecoder(file).Decode(&binaries); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	idCount := make(map[string]int)
	for _, binary := range binaries {
		if len(binary.Entity) <= 1 || len(binary.Id) <= 1 {
			return fmt.Errorf("error while validating binary JSON: %v", binary)
		}

		idCount[binary.Id]++
		if idCount[binary.Id] > 1 {
			return fmt.Errorf("duplicate binary ID: %s", binary.Id)
		}
	}

	return nil
}

func validateContractJSON(file *os.File) error {
	var contracts []yacar_util.Contract

	if err := json.NewDecoder(file).Decode(&contracts); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	idCount := make(map[string]int)
	for _, contract := range contracts {
		if len(contract.Entity) <= 1 || len(contract.Id) <= 1 {
			return fmt.Errorf("error while validating contract JSON: %v", contract)
		}

		idCount[contract.Id]++
		if idCount[contract.Id] > 1 {
			return fmt.Errorf("duplicate contract ID: %s", contract.Id)
		}
	}

	return nil
}

func validateEntityJSON(file *os.File) error {
	var entities []yacar_util.Entity

	if err := json.NewDecoder(file).Decode(&entities); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	stringCount := make(map[string]int)
	for _, entity := range entities {
		if len(entity.Name) <= 1 {
			return fmt.Errorf("error while validating entity JSON: %v", entity)
		}

		stringCount[entity.Name]++
		if stringCount[entity.Name] > 1 {
			return fmt.Errorf("duplicate entity name: %s", entity.Name)
		}
	}

	return nil
}

func validatePoolJSON(file *os.File) error {
	var pools []yacar_util.Pool

	if err := json.NewDecoder(file).Decode(&pools); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	idCount := make(map[string]int)
	for _, pool := range pools {
		if !pool.IsMinimallyPopulated() {
			return fmt.Errorf("error while validating pool JSON: %v", pool)
		}

		idCount[pool.Id]++
		if idCount[pool.Id] > 1 {
			return fmt.Errorf("duplicate pool ID: %s", pool.Id)
		}
	}

	return nil
}
