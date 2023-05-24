package validator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/coinhall/yacar/internal/enums"
	"github.com/coinhall/yacarsdk/v2"
)

func Start(filePaths []string) {
	validateYacarJSONs(filePaths)
	log.Println("Validated JSONs successfully...")
}

func validateYacarJSONs(filePaths []string) {
	var wg sync.WaitGroup
	for _, filePath := range filePaths {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()

			file, err := os.Open(filePath)
			if err != nil {
				log.Fatalf("error while opening file: %s", err)
			}
			defer file.Close()

			if err := validateYacarJSON(filePath, file); err != nil {
				log.Fatalf("%s\npath: %s", err, filePath)
			}

		}(filePath)
	}
	wg.Wait()
}

func validateYacarJSON(filePath string, file *os.File) error {
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
		return fmt.Errorf("unknown file path: %s", filePath)
	}
}

func validateAccountJSON(file *os.File) error {
	var accounts []yacarsdk.Account

	if err := json.NewDecoder(file).Decode(&accounts); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	idCount := make(map[string]struct{})
	for _, account := range accounts {
		if !account.IsMinimallyPopulated() {
			return fmt.Errorf("account ID %s is not minimally populated", account.Id)
		}

		if _, ok := idCount[account.Id]; !ok {
			idCount[account.Id] = struct{}{}
		} else {
			return fmt.Errorf("duplicate account ID: %s", account.Id)
		}
	}
	return nil
}

func validateAssetJSON(file *os.File) error {
	var assets []yacarsdk.Asset

	if err := json.NewDecoder(file).Decode(&assets); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	idCount := make(map[string]struct{})
	for _, asset := range assets {
		if !asset.IsMinimallyPopulated() {
			return fmt.Errorf("asset ID %s is not minimally populated", asset.Id)
		}

		if _, ok := idCount[asset.Id]; !ok {
			idCount[asset.Id] = struct{}{}
		} else {
			return fmt.Errorf("duplicate asset ID: %s", asset.Id)
		}
	}

	return nil
}

func validateBinaryJSON(file *os.File) error {
	var binaries []yacarsdk.Binary

	if err := json.NewDecoder(file).Decode(&binaries); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	idCount := make(map[string]struct{})
	for _, binary := range binaries {
		if !binary.IsMinimallyPopulated() {
			return fmt.Errorf("binary ID %s is not minimally populated", binary.Id)
		}

		if _, ok := idCount[binary.Id]; !ok {
			idCount[binary.Id] = struct{}{}
		} else {
			return fmt.Errorf("duplicate binary ID: %s", binary.Id)
		}
	}

	return nil
}

func validateContractJSON(file *os.File) error {
	var contracts []yacarsdk.Contract

	if err := json.NewDecoder(file).Decode(&contracts); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	idCount := make(map[string]struct{})
	for _, contract := range contracts {
		if !contract.IsMinimallyPopulated() {
			return fmt.Errorf("contract ID %s is not minimally populated", contract.Id)
		}

		if _, ok := idCount[contract.Id]; !ok {
			idCount[contract.Id] = struct{}{}
		} else {
			return fmt.Errorf("duplicate contract ID: %s", contract.Id)
		}
	}

	return nil
}

func validateEntityJSON(file *os.File) error {
	var entities []yacarsdk.Entity

	if err := json.NewDecoder(file).Decode(&entities); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	entityCount := make(map[string]struct{})
	for _, entity := range entities {
		if !entity.IsMinimallyPopulated() {
			return fmt.Errorf("entity name %s is not minimally populated", entity.Name)
		}

		if _, ok := entityCount[entity.Name]; !ok {
			entityCount[entity.Name] = struct{}{}
		} else {
			return fmt.Errorf("duplicate entity name: %s", entity.Name)
		}
	}

	return nil
}

func validatePoolJSON(file *os.File) error {
	var pools []yacarsdk.Pool

	if err := json.NewDecoder(file).Decode(&pools); err != nil {
		return fmt.Errorf("error while decoding JSON: %s", err)
	}

	idCount := make(map[string]struct{})
	for _, pool := range pools {
		if !pool.IsMinimallyPopulated() {
			return fmt.Errorf("pool ID %s is not minimally populated", pool.Id)
		}

		if _, ok := idCount[pool.Id]; !ok {
			idCount[pool.Id] = struct{}{}
		} else {
			return fmt.Errorf("duplicate pool ID: %s", pool.Id)
		}
	}

	return nil
}
