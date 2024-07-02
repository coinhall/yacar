package validator

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/coinhall/yacarsdk/v2"
)

func Start(filePaths []string) {
	validateYacarJSONs(filePaths)
	log.Println("Validated JSONs successfully...")
}

func validateYacarJSONs(filePaths []string) {
	chainFileMap := map[string]map[string]*os.File{}
	for _, fp := range filePaths {
		fp = filepath.ToSlash(fp)
		fpElements := strings.Split(fp, "/")
		chain := fpElements[len(fpElements)-2]
		filetype := strings.Split(fpElements[len(fpElements)-1], ".")[0]
		if _, ok := chainFileMap[chain]; !ok {
			chainFileMap[chain] = map[string]*os.File{}
		}

		file, err := os.Open(fp)
		if err != nil {
			panic(fmt.Sprintf("error while opening file: %s", err))
		}
		chainFileMap[chain][filetype] = file
	}
	defer closeChainFileMaps(chainFileMap)

	if err := validateYacarJSON(chainFileMap); err != nil {
		panic(err)
	}
}

func closeChainFileMaps(cfm map[string]map[string]*os.File) {
	for _, fm := range cfm {
		for _, f := range fm {
			f.Close()
		}
	}
}

func validateYacarJSON(cfm map[string]map[string]*os.File) error {
	for chain, filemap := range cfm {
		for _, file := range filemap {
			var err error
			switch {
			case strings.Contains(file.Name(), "account"):
				err = validateAccountJSON(file)
			case strings.Contains(file.Name(), "asset"):
				entityFile := cfm[chain]["entity"]
				err = validateAssetJSON(file, entityFile)
			case strings.Contains(file.Name(), "binary"):
				err = validateBinaryJSON(file)
			case strings.Contains(file.Name(), "contract"):
				err = validateContractJSON(file)
			case strings.Contains(file.Name(), "entity"):
				accountFile := cfm[chain]["account"]
				assetFile := cfm[chain]["asset"]
				binaryFile := cfm[chain]["binary"]
				contractFile := cfm[chain]["contract"]
				err = validateEntityJSON(file, accountFile, assetFile, binaryFile, contractFile)
			case strings.Contains(file.Name(), "pool"):
				err = validatePoolJSON(file)
			default:
				err = fmt.Errorf("unknown file type: %s", file.Name())
			}

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateAccountJSON(file *os.File) error {
	var accounts []yacarsdk.Account

	file.Seek(0, io.SeekStart)
	if err := json.NewDecoder(file).Decode(&accounts); err != nil {
		return fmt.Errorf("error while decoding account JSON: %s", err)
	}

	_, err := yacarsdk.ValidateAccounts(accounts)
	return err
}

func validateAssetJSON(assetFile, entityFile *os.File) error {
	var (
		assets   []yacarsdk.Asset
		entities []yacarsdk.Entity
	)

	// Reset file offset to beginning of file in case it was read before, not doing so would cause
	// an EOF error when trying to decode the JSON
	assetFile.Seek(0, io.SeekStart)
	if err := json.NewDecoder(assetFile).Decode(&assets); err != nil {
		return fmt.Errorf("error while decoding asset JSON for asset validation: %s", assetFile.Name())
	}

	entityFile.Seek(0, io.SeekStart)
	if err := json.NewDecoder(entityFile).Decode(&entities); err != nil {
		return fmt.Errorf("error while decoding entity JSON for asset validation: %s", err)
	}

	_, err := yacarsdk.ValidateAssets(assets, entities)
	return err
}

func validateBinaryJSON(file *os.File) error {
	var binaries []yacarsdk.Binary

	file.Seek(0, io.SeekStart)
	if err := json.NewDecoder(file).Decode(&binaries); err != nil {
		return fmt.Errorf("error while decoding binary JSON: %s", err)
	}

	_, err := yacarsdk.ValidateBinaries(binaries)
	return err
}

func validateContractJSON(file *os.File) error {
	var contracts []yacarsdk.Contract

	file.Seek(0, io.SeekStart)
	if err := json.NewDecoder(file).Decode(&contracts); err != nil {
		return fmt.Errorf("error while decoding contract JSON: %s", err)
	}

	_, err := yacarsdk.ValidateContracts(contracts)
	return err
}

func validateEntityJSON(entityFile, accountFile, assetFile, binaryFile, contractFile *os.File) error {
	var (
		entities  []yacarsdk.Entity
		accounts  []yacarsdk.Account
		assets    []yacarsdk.Asset
		binaries  []yacarsdk.Binary
		contracts []yacarsdk.Contract
	)

	// Reset file offset to beginning of file in case it was read before, not doing so would cause
	// an EOF error when trying to decode the JSON
	entityFile.Seek(0, io.SeekStart)
	if err := json.NewDecoder(entityFile).Decode(&entities); err != nil {
		return fmt.Errorf("error while decoding entity JSON for entity validation: %s", err)
	}

	usedEntities := map[string]struct{}{}

	accountFile.Seek(0, io.SeekStart)
	if err := json.NewDecoder(accountFile).Decode(&accounts); err != nil {
		return fmt.Errorf("error while decoding account JSON for entity validation: %s", err)
	}
	for _, account := range accounts {
		usedEntities[account.Entity] = struct{}{}
	}

	assetFile.Seek(0, io.SeekStart)
	if err := json.NewDecoder(assetFile).Decode(&assets); err != nil {
		return fmt.Errorf("error while decoding asset JSON for entity validation: %s", err)
	}
	for _, asset := range assets {
		usedEntities[asset.Entity] = struct{}{}
	}

	binaryFile.Seek(0, io.SeekStart)
	if err := json.NewDecoder(binaryFile).Decode(&binaries); err != nil {
		return fmt.Errorf("error while decoding binary JSON for entity validation: %s", err)
	}
	for _, binary := range binaries {
		usedEntities[binary.Entity] = struct{}{}
	}

	contractFile.Seek(0, io.SeekStart)
	if err := json.NewDecoder(contractFile).Decode(&contracts); err != nil {
		return fmt.Errorf("error while decoding contract JSON for entity validation: %s", err)
	}
	for _, contract := range contracts {
		usedEntities[contract.Entity] = struct{}{}
	}

	_, err := yacarsdk.ValidateEntities(entities, usedEntities)
	return err
}

func validatePoolJSON(file *os.File) error {
	var pools []yacarsdk.Pool

	if err := json.NewDecoder(file).Decode(&pools); err != nil {
		return fmt.Errorf("error while decoding pool JSON: %s", err)
	}

	_, err := yacarsdk.ValidatePools(pools)
	return err
}
