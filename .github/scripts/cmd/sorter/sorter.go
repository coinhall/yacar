package sorter

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/coinhall/yacarsdk/v2"
)

type CustomAsset struct {
	Id             string `json:"id"`
	Entity         string `json:"entity,omitempty"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Decimals       string `json:"decimals"`
	Type           string `json:"type"`
	CircSupplyAPI  string `json:"-"`
	TotalSupplyAPI string `json:"-"`
	Icon           string `json:"icon,omitempty"`
	CoinMarketCap  string `json:"coinmarketcap,omitempty"`
	CoinGecko      string `json:"coingecko,omitempty"`
}

func Start(filePaths []string) {
	sortYacarJSONs(filePaths)
	log.Println("Sorted JSONs successfully...")
}

func sortYacarJSONs(filePaths []string) {
	var wg sync.WaitGroup
	for _, filePath := range filePaths {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()

			file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
			if err != nil {
				log.Panicf("error while opening file: %s", err)
			}
			defer file.Close()

			sortedJSON := sortYacarJSON(file)
			writeYacarJSON(file, sortedJSON)

		}(filePath)
	}
	wg.Wait()
}

func sortYacarJSON(file *os.File) interface{} {
	switch {
	case strings.Contains(file.Name(), "account"):
		return sortAccountJSON(file)
	case strings.Contains(file.Name(), "asset"):
		return sortAssetJSON(file)
	case strings.Contains(file.Name(), "binary"):
		return sortBinaryJSON(file)
	case strings.Contains(file.Name(), "contract"):
		return sortContractJSON(file)
	case strings.Contains(file.Name(), "entity"):
		return sortEntityJSON(file)
	case strings.Contains(file.Name(), "pool"):
		return sortPoolJSON(file)
	default:
		log.Panic("unable to sort unknown JSON type")
		return nil
	}
}

func sortAccountJSON(file *os.File) interface{} {
	var accounts []yacarsdk.Account

	if err := json.NewDecoder(file).Decode(&accounts); err != nil {
		log.Panicf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedAccountOrder(accounts))

	return accounts
}

func sortAssetJSON(file *os.File) interface{} {
	var assets []yacarsdk.Asset
	var customAssets []CustomAsset

	if err := json.NewDecoder(file).Decode(&customAssets); err != nil {
		log.Panicf("error while decoding JSON: %s", err)
	}

	for _, customAsset := range customAssets {
		assets = append(assets, yacarsdk.Asset{
			Id:             customAsset.Id,
			Entity:         customAsset.Entity,
			Name:           customAsset.Name,
			Symbol:         customAsset.Symbol,
			Decimals:       customAsset.Decimals,
			Type:           customAsset.Type,
			CircSupplyAPI:  "",
			TotalSupplyAPI: "",
			Icon:           customAsset.Icon,
			CoinMarketCap:  customAsset.CoinMarketCap,
			CoinGecko:      customAsset.CoinGecko,
		})
	}

	sort.Stable(yacarsdk.ByEnforcedAssetOrder(assets))

	return assets
}

func sortBinaryJSON(file *os.File) interface{} {
	var binaries []yacarsdk.Binary

	if err := json.NewDecoder(file).Decode(&binaries); err != nil {
		log.Panicf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedBinaryOrder(binaries))

	return binaries
}

func sortContractJSON(file *os.File) interface{} {
	var contracts []yacarsdk.Contract

	if err := json.NewDecoder(file).Decode(&contracts); err != nil {
		log.Panicf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedContractOrder(contracts))

	return contracts
}

func sortEntityJSON(file *os.File) interface{} {
	var entities []yacarsdk.Entity

	if err := json.NewDecoder(file).Decode(&entities); err != nil {
		log.Panicf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedEntityOrder(entities))

	return entities
}

func sortPoolJSON(file *os.File) interface{} {
	var pools []yacarsdk.Pool

	if err := json.NewDecoder(file).Decode(&pools); err != nil {
		log.Panicf("error while decoding JSON: %s", err)
	}

	sort.Stable(yacarsdk.ByEnforcedPoolOrder(pools))

	return pools
}

func writeYacarJSON(file *os.File, data interface{}) {
	// Clear file
	if err := file.Truncate(0); err != nil {
		log.Panicf("error while truncating file: %s", err)
	}

	// Move cursor to the beginning of the file
	if _, err := file.Seek(0, 0); err != nil {
		log.Panicf("error while seeking file: %s", err)
	}

	enc := json.NewEncoder(file)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		log.Panicf("error while encoding JSON: %s", err)
	}
}
