package validator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

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
				log.Panicf("error while opening file: %s", err)
			}
			defer file.Close()

			if err := validateYacarJSON(file); err != nil {
				log.Panicf("%s\npath: %s", err, filePath)
			}

		}(filePath)
	}
	wg.Wait()
}

func validateYacarJSON(file *os.File) error {
	switch {
	case strings.Contains(file.Name(), "account"):
		return validateAccountJSON(file)
	case strings.Contains(file.Name(), "asset"):
		return validateAssetJSON(file)
	case strings.Contains(file.Name(), "binary"):
		return validateBinaryJSON(file)
	case strings.Contains(file.Name(), "contract"):
		return validateContractJSON(file)
	case strings.Contains(file.Name(), "entity"):
		return validateEntityJSON(file)
	case strings.Contains(file.Name(), "pool"):
		return validatePoolJSON(file)
	}
	return nil
}

func validateAccountJSON(file *os.File) error {
	var accounts []yacarsdk.Account

	if err := json.NewDecoder(file).Decode(&accounts); err != nil {
		return fmt.Errorf("error while decoding account JSON: %s", err)
	}

	for _, account := range accounts {
		if !account.IsMinimallyPopulated() {
			return fmt.Errorf("account ID %s is not minimally populated", account.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, account := range accounts {
		if _, ok := idCount[account.Id]; ok {
			return fmt.Errorf("duplicate account ID: %s", account.Id)
		}

		idCount[account.Id] = struct{}{}
	}

	return nil
}

func validateAssetJSON(file *os.File) error {
	var assets []yacarsdk.Asset

	if err := json.NewDecoder(file).Decode(&assets); err != nil {
		return fmt.Errorf("error while decoding asset JSON: %s", err)
	}

	for _, asset := range assets {
		if !asset.IsMinimallyPopulated() {
			return fmt.Errorf("asset ID %s is not minimally populated", asset.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, asset := range assets {
		if _, ok := idCount[asset.Id]; ok {
			return fmt.Errorf("duplicate asset ID: %s", asset.Id)
		}
		idCount[asset.Id] = struct{}{}
	}

	txCheck := make(map[string]struct{})
	for _, asset := range assets {
		if _, ok := txCheck[asset.VerificationTx]; ok {
			return fmt.Errorf("duplicate asset tx hash: %s", asset.Id)
		}
		txCheck[asset.VerificationTx] = struct{}{}
	}
	return nil
}

func validateBinaryJSON(file *os.File) error {
	var binaries []yacarsdk.Binary

	if err := json.NewDecoder(file).Decode(&binaries); err != nil {
		return fmt.Errorf("error while decoding binary JSON: %s", err)
	}

	for _, binary := range binaries {
		if !binary.IsMinimallyPopulated() {
			return fmt.Errorf("binary ID %s is not minimally populated", binary.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, binary := range binaries {
		if _, ok := idCount[binary.Id]; ok {
			return fmt.Errorf("duplicate binary ID: %s", binary.Id)
		}

		idCount[binary.Id] = struct{}{}
	}

	return nil
}

func validateContractJSON(file *os.File) error {
	var contracts []yacarsdk.Contract

	if err := json.NewDecoder(file).Decode(&contracts); err != nil {
		return fmt.Errorf("error while decoding contract JSON: %s", err)
	}

	for _, contract := range contracts {
		if !contract.IsMinimallyPopulated() {
			return fmt.Errorf("contract ID %s is not minimally populated", contract.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, contract := range contracts {
		if _, ok := idCount[contract.Id]; ok {
			return fmt.Errorf("duplicate contract ID: %s", contract.Id)
		}

		idCount[contract.Id] = struct{}{}
	}

	return nil
}

func validateEntityJSON(file *os.File) error {
	var entities []yacarsdk.Entity

	if err := json.NewDecoder(file).Decode(&entities); err != nil {
		return fmt.Errorf("error while decoding entity JSON: %s", err)
	}

	for _, entity := range entities {
		if !entity.IsMinimallyPopulated() {
			return fmt.Errorf("entity name %s is not minimally populated", entity.Name)
		}
	}

	entityCount := make(map[string]struct{})
	for _, entity := range entities {
		if _, ok := entityCount[entity.Name]; ok {
			return fmt.Errorf("duplicate entity name: %s", entity.Name)
		}

		entityCount[entity.Name] = struct{}{}
	}

	return nil
}

func validatePoolJSON(file *os.File) error {
	var pools []yacarsdk.Pool

	if err := json.NewDecoder(file).Decode(&pools); err != nil {
		return fmt.Errorf("error while decoding pool JSON: %s", err)
	}

	for _, pool := range pools {
		if !pool.IsMinimallyPopulated() {
			return fmt.Errorf("pool ID %s is not minimally populated", pool.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, pool := range pools {
		if _, ok := idCount[pool.Id]; ok {
			return fmt.Errorf("duplicate pool ID: %s", pool.Id)
		}

		idCount[pool.Id] = struct{}{}
	}

	return nil
}
