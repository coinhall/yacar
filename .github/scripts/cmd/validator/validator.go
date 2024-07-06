package validator

import (
	"log"
	"path/filepath"
	"slices"
	"strings"

	"github.com/coinhall/yacar/internal/unmarshaler"
	"github.com/coinhall/yacar/internal/walker"
	"github.com/coinhall/yacar/internal/yacar"
	"github.com/coinhall/yacarsdk/v2"
	"golang.org/x/exp/maps"
)

func Start(filePaths []string, ignoreErrors map[string]struct{}) {
	validateYacarJSONs(filePaths, ignoreErrors)
	log.Println("Validated JSONs successfully...")
}

func validateYacarJSONs(filePaths []string, ignore map[string]struct{}) {
	crm := getChainResources(filePaths)
	keys := maps.Keys(crm)
	slices.Sort(keys)
	for _, k := range keys {
		var (
			asset, entity, pool, contract, binary, account string
		)
		for _, fp := range crm[k] {
			switch yacar.MustParse(walker.GetFileNameNoSuffix(fp, yacar.FileSuffix)) {
			case yacar.Asset:
				asset = fp
			case yacar.Entity:
				entity = fp
			case yacar.Pool:
				pool = fp
			case yacar.Contract:
				contract = fp
			case yacar.Binary:
				binary = fp
			case yacar.Account:
				account = fp
			default:
				panic("unhandled case")
			}
		}

		errorCheck(handleAccount(account), ignore)
		errorCheck(handlePool(pool), ignore)
		errorCheck(handleContract(contract), ignore)
		errorCheck(handleBinary(binary), ignore)
		errorCheck(handleAsset(asset, entity), ignore)
		errorCheck(handleEntity(entity, account, asset, binary,
			contract, pool), ignore)
	}
}

func getChainResources(filePaths []string) map[string][]string {
	crm := map[string][]string{}
	for _, fp := range filePaths {
		fp = filepath.ToSlash(fp)
		split := strings.Split(filepath.Dir(fp), "/")
		chain := split[len(split)-1]
		r, ok := crm[chain]
		if !ok {
			r = make([]string, 0)
		}

		r = append(r, fp)
		crm[chain] = r
	}
	return crm
}

func handleAccount(fp string) error {
	accounts, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Account, 0))
	if err != nil {
		return err
	}

	_, err = yacarsdk.ValidateAccounts(accounts)
	return err
}

func handleAsset(assetFile, entityFile string) error {
	assets, err := unmarshaler.UnmarshalInto(assetFile, make([]yacarsdk.Asset, 0))
	if err != nil {
		return err
	}

	entities, err := unmarshaler.UnmarshalInto(entityFile, make([]yacarsdk.Entity, 0))
	if err != nil {
		return err
	}

	_, err = yacarsdk.ValidateAssets(assets, entities)
	return err
}

func handleBinary(fp string) error {
	if len(fp) == 0 {
		return nil
	}

	binaries, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Binary, 0))
	if err != nil {
		return err
	}

	_, err = yacarsdk.ValidateBinaries(binaries)
	return err
}

func handleContract(fp string) error {
	if len(fp) == 0 {
		return nil
	}

	contracts, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Contract, 0))
	if err != nil {
		panic(err)
	}

	_, err = yacarsdk.ValidateContracts(contracts)
	return err
}

func handleEntity(entity, account, asset, binary, contract, pool string) error {
	entities, err := unmarshaler.UnmarshalInto(entity, make([]yacarsdk.Entity, 0))
	if err != nil {
		return err
	}
	usedEntities := map[string]struct{}{}

	if len(account) > 0 {
		accounts, err := unmarshaler.UnmarshalInto(account, make([]yacarsdk.Account, 0))
		if err != nil {
			return err
		}
		for _, a := range accounts {
			usedEntities[a.Entity] = struct{}{}
		}
	}

	if len(asset) > 0 {
		assets, err := unmarshaler.UnmarshalInto(asset, make([]yacarsdk.Asset, 0))
		if err != nil {
			return err
		}
		for _, a := range assets {
			usedEntities[a.Entity] = struct{}{}
		}
	}

	if len(binary) > 0 {
		binaries, err := unmarshaler.UnmarshalInto(binary, make([]yacarsdk.Binary, 0))
		if err != nil {
			return err
		}
		for _, a := range binaries {
			usedEntities[a.Entity] = struct{}{}
		}
	}

	if len(contract) > 0 {
		contracts, err := unmarshaler.UnmarshalInto(contract, make([]yacarsdk.Contract, 0))
		if err != nil {
			return err
		}
		for _, a := range contracts {
			usedEntities[a.Entity] = struct{}{}
		}
	}

	_, err = yacarsdk.ValidateEntities(entities, usedEntities)
	return err
}

func handlePool(fp string) error {
	pools, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Pool, 0))
	if err != nil {
		return err
	}

	_, err = yacarsdk.ValidatePools(pools)
	return err
}

func errorCheck(err error, ignore map[string]struct{}) {
	if err == nil {
		return
	}

	if _, ok := ignore[err.Error()]; !ok {
		log.Fatalln(err)
	}
}
