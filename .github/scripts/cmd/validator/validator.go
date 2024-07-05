package validator

import (
	"cmp"
	"log"
	"path/filepath"
	"strings"

	"github.com/coinhall/yacar/internal/unmarshaler"
	"github.com/coinhall/yacar/internal/walker"
	"github.com/coinhall/yacarsdk/v2"
)

func Start(filePaths []string) {
	validateYacarJSONs(filePaths)
	log.Println("Validated JSONs successfully...")
}

type resources struct {
	account, asset, binary, contract, entity, pool string
}

func validateYacarJSONs(filePaths []string) {
	var err error
	crm := getChainResources(filePaths)
	for _, res := range crm {
		path := cmp.Or(res.account, res.asset, res.binary, res.contract, res.entity, res.pool)
		switch walker.MustParse(walker.TrimJsonSuffixFromFullPath(path)) {
		case walker.Account:
			err = handleAccount(res.account)
		case walker.Asset:
			err = handleAsset(res.asset, res.entity)
		case walker.Binary:
			err = handleBinary(res.binary)
		case walker.Contract:
			err = handleContract(res.contract)
		case walker.Entity:
			err = handleEntity(res)
		case walker.Pool:
			err = handlePool(res.pool)
		default:
			panic("unhandled case")
		}

		if err != nil {
			panic(err)
		}
	}
}

func getChainResources(filePaths []string) map[string]*resources {
	crm := map[string]*resources{}
	for _, fp := range filePaths {
		fp = filepath.ToSlash(fp)
		split := strings.Split(filepath.Dir(fp), "/")
		chain := split[len(split)-1]
		r, ok := crm[chain]
		if !ok {
			r = &resources{}
		}
		switch walker.MustParse(walker.TrimJsonSuffixFromFullPath(fp)) {
		case walker.Account:
			r.account = fp
		case walker.Asset:
			r.asset = fp
		case walker.Binary:
			r.binary = fp
		case walker.Contract:
			r.contract = fp
		case walker.Entity:
			r.entity = fp
		case walker.Pool:
			r.pool = fp
		default:
			panic("unhandled case")
		}
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

// Optional
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

// Optional
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

func handleEntity(res *resources) error {
	entities, err := unmarshaler.UnmarshalInto(res.entity, make([]yacarsdk.Entity, 0))
	if err != nil {
		return err
	}
	usedEntities := map[string]struct{}{}

	if len(res.account) > 0 {
		accounts, err := unmarshaler.UnmarshalInto(res.account, make([]yacarsdk.Account, 0))
		if err != nil {
			return err
		}
		for _, a := range accounts {
			usedEntities[a.Entity] = struct{}{}
		}
	}

	if len(res.asset) > 0 {
		assets, err := unmarshaler.UnmarshalInto(res.asset, make([]yacarsdk.Asset, 0))
		if err != nil {
			return err
		}
		for _, a := range assets {
			usedEntities[a.Entity] = struct{}{}
		}
	}

	if len(res.binary) > 0 {
		binaries, err := unmarshaler.UnmarshalInto(res.binary, make([]yacarsdk.Binary, 0))
		if err != nil {
			return err
		}
		for _, a := range binaries {
			usedEntities[a.Entity] = struct{}{}
		}
	}

	if len(res.contract) > 0 {
		contracts, err := unmarshaler.UnmarshalInto(res.contract, make([]yacarsdk.Contract, 0))
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
