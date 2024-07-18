package ibcpropagator

import (
	"fmt"
	"log"
	"path/filepath"
	"slices"
	"strings"

	"github.com/coinhall/yacar/internal/unmarshaler"
	"github.com/coinhall/yacar/internal/writer"
	"github.com/coinhall/yacarsdk/v2"
)

func Start(filePaths []string) {
	chainPaths, chainAssetsMap := getChainPathsAssets(filePaths)
	updatedAssetMap := resolveBackwards(chainAssetsMap)
	updateFiles(chainPaths, updatedAssetMap)
	log.Println("Propogated native assets successfully...")
}

func getChainPathsAssets(filePaths []string) (map[string]string, map[string][]yacarsdk.Asset) {
	chainPaths := make(map[string]string)
	chainAssetsMap := make(map[string][]yacarsdk.Asset)
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

		chainAssetsMap[chain] = append(chainAssetsMap[chain], assets...)
	}

	return chainPaths, chainAssetsMap
}

func resolveBackwards(chainAssetsMap map[string][]yacarsdk.Asset) map[string][]yacarsdk.Asset {
	for chain, assets := range chainAssetsMap {
		for i, asset := range assets {
			// Skip if asset is not IBC or does not yet have an associated origin
			if asset.Type != "ibc" || len(asset.OriginId) == 0 {
				continue
			}

			rootAssets, ok := chainAssetsMap[asset.OriginChain]
			if !ok {
				log.Println("origin chain not in YACAR, skipping propagation... -", chain, asset.Id)
				continue
			}

			isMatchingRootAsset := func(ra yacarsdk.Asset) bool { return ra.Id == asset.OriginId }
			ri := slices.IndexFunc(rootAssets, isMatchingRootAsset)
			if ri == -1 {
				log.Fatalln("missing root asset -", asset.OriginId)
			}
			rootAsset := rootAssets[ri]
			asset.Name = rootAsset.Name
			asset.Symbol = rootAsset.Symbol
			asset.Icon = rootAsset.Icon

			// Update asset
			assets[i] = asset
		}

		// Update chain assets
		chainAssetsMap[chain] = assets
	}

	return chainAssetsMap
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
