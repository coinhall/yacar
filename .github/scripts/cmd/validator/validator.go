package validator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
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
		for filetype, file := range filemap {
			log.Printf("Validating %s %s JSON...", chain, filetype)
			var err error
			switch {
			case strings.Contains(file.Name(), "account"):
				err = validateAccountJSON(file)
			case strings.Contains(file.Name(), "asset"):
				err = validateAssetJSON(file, cfm[chain]["entity"])
			case strings.Contains(file.Name(), "binary"):
				err = validateBinaryJSON(file)
			case strings.Contains(file.Name(), "contract"):
				err = validateContractJSON(file)
			case strings.Contains(file.Name(), "entity"):
				err = validateEntityJSON(file, cfm[chain]["asset"])
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

func validateAssetJSON(assetFile, entityFile *os.File) error {
	var (
		assets   []yacarsdk.Asset
		entities []yacarsdk.Entity
	)

	if err := json.NewDecoder(assetFile).Decode(&assets); err != nil {
		return fmt.Errorf("error while decoding asset JSON for asset validation: %s", assetFile.Name())
	}

	if err := json.NewDecoder(entityFile).Decode(&entities); err != nil {
		return fmt.Errorf("error while decoding entity JSON for asset validation: %s", err)
	}

	for _, asset := range assets {
		if !asset.IsMinimallyPopulated() {
			return fmt.Errorf("asset ID %s is not minimally populated", asset.Id)
		}

		if asset.Id == asset.Name {
			return fmt.Errorf("asset name for %s cannot be the asset ID", asset.Id)
		}

		if asset.Id == asset.Symbol {
			return fmt.Errorf("asset symbol for %s cannot be the asset ID", asset.Id)
		}

		if len(asset.Symbol) > 20 {
			return fmt.Errorf("asset symbol for %s cannot be longer than 20 characters", asset.Id)
		}
	}

	idCount := make(map[string]struct{})
	for _, asset := range assets {
		if _, ok := idCount[asset.Id]; ok {
			return fmt.Errorf("duplicate asset ID: %s", asset.Id)
		}
		idCount[asset.Id] = struct{}{}
	}

	// Circ supply check
	for _, asset := range assets {
		if len(asset.CircSupply) > 0 && len(asset.CircSupplyAPI) > 0 {
			return fmt.Errorf("[%s] either 'circ_supply' or 'circ_supply_api' must be specified, but not both", asset.Id)
		}

		if len(asset.CircSupply) > 0 {
			if parsed, err := strconv.ParseFloat(asset.CircSupply, 64); err != nil && parsed > 0 {
				return fmt.Errorf("[%s] 'circ_supply' must be float greater than 0", asset.Id)
			}
		}
	}

	// Total supply check
	for _, asset := range assets {
		if len(asset.TotalSupply) > 0 && len(asset.TotalSupplyAPI) > 0 {
			return fmt.Errorf("[%s] either 'total_supply' or 'total_supply_api' must be specified, but not both", asset.Id)
		}

		if len(asset.TotalSupply) > 0 {
			if parsed, err := strconv.ParseFloat(asset.TotalSupply, 64); err != nil && parsed > 0 {
				return fmt.Errorf("[%s] 'total_supply' must be number greater than 0", asset.Id)
			}
		}
	}

	// Corresponding entity check
	entityNameSet := map[string]struct{}{}
	for _, entity := range entities {
		entityNameSet[entity.Name] = struct{}{}
	}
	for _, asset := range assets {
		if asset.Entity == "" {
			continue
		}
		if _, ok := entityNameSet[asset.Entity]; ok {
			continue
		}

		return fmt.Errorf("[%s] entity %s does not exists", asset.Id, asset.Entity)
	}

	// Non-permissioned DEX TxHash must be unique
	permissionedDex := map[string]struct{}{
		"osmosis-main": {},
		"kujira-fin":   {},
	}
	txCheck := make(map[string]struct{})
	for _, asset := range assets {
		// If asset is from a permissioned DEX or is empty, skip
		if _, ok := permissionedDex[asset.VerificationTx]; ok || asset.VerificationTx == "" {
			continue
		}

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

func validateEntityJSON(entityFile, assetFile *os.File) error {
	var (
		entities []yacarsdk.Entity
		assets   []yacarsdk.Asset
	)

	if err := json.NewDecoder(entityFile).Decode(&entities); err != nil {
		return fmt.Errorf("error while decoding entity JSON for entity validation: %s", err)
	}

	if err := json.NewDecoder(assetFile).Decode(&assets); err != nil {
		return fmt.Errorf("error while decoding asset JSON for entity validation: %s", assetFile.Name())
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

	// // Corresponding asset check
	// assetEntityNameSet := map[string]struct{}{}
	// for _, asset := range assets {
	// 	if asset.Entity == "" {
	// 		continue
	// 	}
	// 	assetEntityNameSet[asset.Entity] = struct{}{}
	// }
	// for _, entity := range entities {
	// 	if _, ok := assetEntityNameSet[entity.Name]; ok {
	// 		continue
	// 	}

	// 	return fmt.Errorf("entity %s does not correspond to any asset", entity.Name)
	// }

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
