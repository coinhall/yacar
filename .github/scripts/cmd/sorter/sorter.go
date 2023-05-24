package sorter

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/coinhall/yacar/internal/enums"
	"github.com/coinhall/yacarsdk/v2"
)

func Start(filePaths []string) {
	sortedJSONsMap := sortYacarJSONs(filePaths)
	writeYacarJSONs(sortedJSONsMap)
	log.Println("Sorted JSONs successfully...")
}

func sortYacarJSONs(filePaths []string) map[string]interface{} {
	var wg sync.WaitGroup
	var mu sync.Mutex
	sm := make(map[string]interface{}, len(filePaths))

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

			mu.Lock()
			sm[filePath] = sorted
			mu.Unlock()
		}(filePath)
	}
	wg.Wait()

	return sm
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
	var accounts []yacarsdk.Account

	if err := json.NewDecoder(file).Decode(&accounts); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedAccountOrder(accounts))

	return accounts
}

func sortAssetJSON(file *os.File) interface{} {
	var assets []yacarsdk.Asset

	if err := json.NewDecoder(file).Decode(&assets); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedAssetOrder(assets))

	return assets
}

func sortBinaryJSON(file *os.File) interface{} {
	var binaries []yacarsdk.Binary

	if err := json.NewDecoder(file).Decode(&binaries); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedBinaryOrder(binaries))

	return binaries
}

func sortContractJSON(file *os.File) interface{} {
	var contracts []yacarsdk.Contract

	if err := json.NewDecoder(file).Decode(&contracts); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedContractOrder(contracts))

	return contracts
}

func sortEntityJSON(file *os.File) interface{} {
	var entities []yacarsdk.Entity

	if err := json.NewDecoder(file).Decode(&entities); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedEntityOrder(entities))

	return entities
}

func sortPoolJSON(file *os.File) interface{} {
	var pools []yacarsdk.Pool

	if err := json.NewDecoder(file).Decode(&pools); err != nil {
		log.Fatalf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedPoolOrder(pools))

	return pools
}

func writeYacarJSONs(sortedJSONsMap map[string]interface{}) {

	var wg sync.WaitGroup
	for fp, data := range sortedJSONsMap {
		wg.Add(1)
		go func(fp string, data interface{}) {
			defer wg.Done()

			file, err := os.OpenFile(fp, os.O_RDWR|os.O_TRUNC, 0644)
			if err != nil {
				log.Fatalf("error while opening file: %s", err)
			}
			defer file.Close()

			var sb strings.Builder
			enc := json.NewEncoder(&sb)
			enc.SetEscapeHTML(false)
			enc.SetIndent("", "  ")
			if err := enc.Encode(data); err != nil {
				log.Fatalf("error while encoding JSON: %s", err)
			}

			if err := file.Truncate(0); err != nil {
				log.Fatalf("error while truncating file: %s", err)
			}

			if _, err := file.WriteString(sb.String()); err != nil {
				log.Fatalf("error while writing to file: %s", err)
			}
		}(fp, data)
	}
	wg.Wait()
}
