package ibcpropagator

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/coinhall/yacar/internal/unmarshaler"
	"github.com/coinhall/yacar/internal/writer"
	"github.com/coinhall/yacarsdk/v2"
)

func Start(filePaths []string) {
	chainPaths, chainAssetMap := getChainPathsAssets(filePaths)
	updatedAssetMap := propogateNatives(chainAssetMap)
	updateFiles(chainPaths, updatedAssetMap)
	log.Println("Propogated native assets successfully...")
}

func getChainPathsAssets(filePaths []string) (map[string]string, map[string][]yacarsdk.Asset) {
	chainPaths := make(map[string]string)
	chainAssetMap := make(map[string][]yacarsdk.Asset)
	for _, fp := range filePaths {
		base := filepath.Base(fp)
		if base != "asset.json" {
			continue
		}

		split := strings.Split(filepath.Dir(fp), "/")
		chain := split[len(split)-1]

		chainPaths[chain] = fp

		assets, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Asset, 0))
		if err != nil {
			panic(err)
		}

		chainAssetMap[chain] = append(chainAssetMap[chain], assets...)
	}

	return chainPaths, chainAssetMap
}

func propogateNatives(chainAssetMap map[string][]yacarsdk.Asset) map[string][]yacarsdk.Asset {
	for curChain, curAssets := range chainAssetMap {
		for otherChain, otherAssets := range chainAssetMap {
			if curChain == otherChain {
				continue
			}

			chainAssetMap[otherChain] = modifyDownstreams(curAssets, otherAssets)
		}
	}
	return chainAssetMap
}

func modifyDownstreams(root, downstream []yacarsdk.Asset) []yacarsdk.Asset {
	for _, ra := range root {
		if ra.Type == "ibc" {
			continue
		}

		for i, da := range downstream {
			if da.Type != "ibc" || len(da.OriginId) == 0 {
				continue
			}

			// This might not be unique enough
			if da.OriginId != ra.Id {
				continue
			}

			da.Name = ra.Name
			da.Symbol = ra.Symbol
			downstream[i] = da
		}
	}

	return downstream
}

func updateFiles(chainPaths map[string]string, updatedAssets map[string][]yacarsdk.Asset) {
	for k := range chainPaths {
		if _, ok := updatedAssets[k]; !ok {
			panic(fmt.Sprintf("chain %s not found in updated assets", k))
		}
	}
	for k := range updatedAssets {
		if _, ok := chainPaths[k]; !ok {
			panic(fmt.Sprintf("chain %s not found in chain paths", k))
		}
	}

	for chain, fp := range chainPaths {
		assets, ok := updatedAssets[chain]
		if !ok {
			panic(fmt.Sprintf("chain %s not found in updated assets", chain))
		}

		if err := writer.WriteFile(fp, assets); err != nil {
			panic(err)
		}
	}
}
