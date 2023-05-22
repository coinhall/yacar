package sorter

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/coinhall/yacar/internal/enums"
	"github.com/coinhall/yacar_util"
)

func Start(filePaths []string) {
	// Sort JSONs
	sortedJSONsMap, err := sortYacarJSONs(filePaths)
	if err != nil {
		log.Fatal(err)
	}

	// Write changes to disk, map is guaranteed to be populated
	writeYacarJSONs(sortedJSONsMap)
}

func sortYacarJSONs(filePaths []string) (sync.Map, error) {
	var wg sync.WaitGroup
	var sm sync.Map

	for _, filePath := range filePaths {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()

			file, err := os.Open(filePath)
			if err != nil {
				log.Fatalf("error while opening file: %s", err)
			}
			defer file.Close()

			sorted := sortYacarJSON(filePath, file)
			sm.Store(filePath, sorted)

		}(filePath)
	}
	wg.Wait()

	return sm, nil
}

func sortYacarJSON(filePath string, file *os.File) interface{} {
	switch {
	case strings.Contains(filePath, enums.Account.Name()):
		return sortAccountJSON(file)
	case strings.Contains(filePath, enums.Asset.Name()):
		return sortAssetJSON(file)
	case strings.Contains(filePath, enums.Binary.Name()):
		return sortBinaryJSON(file)
	case strings.Contains(filePath, enums.Contract.Name()):
		return sortContractJSON(file)
	case strings.Contains(filePath, enums.Entity.Name()):
		return sortEntityJSON(file)
	case strings.Contains(filePath, enums.Pool.Name()):
		return sortPoolJSON(file)
	default:
		log.Fatal("unable to sort unknown JSON type")
		return nil
	}
}

func sortAccountJSON(file *os.File) interface{} {
	var accounts []yacar_util.Account

	if err := json.NewDecoder(file).Decode(&accounts); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacar_util.ByEnforcedAccountOrder(accounts))

	return accounts
}

func sortAssetJSON(file *os.File) interface{} {
	var assets []yacar_util.Asset

	if err := json.NewDecoder(file).Decode(&assets); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacar_util.ByEnforcedAssetOrder(assets))

	return assets
}

func sortBinaryJSON(file *os.File) interface{} {
	var binaries []yacar_util.Binary

	if err := json.NewDecoder(file).Decode(&binaries); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacar_util.ByEnforcedBinaryOrder(binaries))

	return binaries
}

func sortContractJSON(file *os.File) interface{} {
	var contracts []yacar_util.Contract

	if err := json.NewDecoder(file).Decode(&contracts); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacar_util.ByEnforcedContractOrder(contracts))

	return contracts
}

func sortEntityJSON(file *os.File) interface{} {
	var entities []yacar_util.Entity

	if err := json.NewDecoder(file).Decode(&entities); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacar_util.ByEnforcedEntityOrder(entities))

	return entities
}

func sortPoolJSON(file *os.File) interface{} {
	var pools []yacar_util.Pool

	if err := json.NewDecoder(file).Decode(&pools); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacar_util.ByEnforcedPoolOrder(pools))

	return pools
}

func writeYacarJSONs(orderedJSONs sync.Map) {
	orderedJSONs.Range(func(key, value interface{}) bool {
		filePath := key.(string)
		orderedJSON := value.([]interface{})

		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalf("error while opening file: %s", err)
		}
		defer file.Close()

		var sb strings.Builder
		jsonEncoder := json.NewEncoder(&sb)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.SetIndent("", "  ")
		if err := jsonEncoder.Encode(orderedJSON); err != nil {
			log.Fatalf("error while encoding JSON: %s", err)
		}

		file.Truncate(0)
		file.WriteString(sb.String())

		return true
	})
}
