package sorter

import (
	"log"
	"path/filepath"
	"sort"

	"github.com/coinhall/yacar/internal/unmarshaler"
	"github.com/coinhall/yacar/internal/walker"
	"github.com/coinhall/yacar/internal/writer"
	"github.com/coinhall/yacar/internal/yacar"
	"github.com/coinhall/yacarsdk/v2"
)

func Start(filePaths []string) {
	sortYacarJSONs(filePaths)
	log.Println("Sorted JSONs successfully...")
}

func sortYacarJSONs(filePaths []string) {
	for _, fp := range filePaths {
		fp = filepath.ToSlash(fp)
		switch yacar.MustParse(walker.GetFileNameNoSuffix(fp, yacar.FileSuffix)) {
		case yacar.Account:
			handleAccount(fp)
		case yacar.Asset:
			handleAsset(fp)
		case yacar.Binary:
			handleBinary(fp)
		case yacar.Contract:
			handleContract(fp)
		case yacar.Entity:
			handleEntity(fp)
		case yacar.Pool:
			handlePool(fp)
		default:
			panic("unhandled case")
		}
	}
}

func handleAccount(fp string) {
	accounts, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Account, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedAccountOrder(accounts))
	if err := writer.WriteFile(fp, accounts); err != nil {
		panic(err)
	}
}

func handleAsset(fp string) {
	assets, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Asset, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedAssetOrder(assets))
	if err := writer.WriteFile(fp, assets); err != nil {
		panic(err)
	}
}

func handleBinary(fp string) {
	binaries, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Binary, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedBinaryOrder(binaries))
	if err := writer.WriteFile(fp, binaries); err != nil {
		panic(err)
	}
}

func handleContract(fp string) {
	contracts, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Contract, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedContractOrder(contracts))
	if err := writer.WriteFile(fp, contracts); err != nil {
		panic(err)
	}
}

func handleEntity(fp string) {
	entities, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Entity, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedEntityOrder(entities))
	if err := writer.WriteFile(fp, entities); err != nil {
		panic(err)
	}
}

func handlePool(fp string) {
	pools, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Pool, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedPoolOrder(pools))
	if err := writer.WriteFile(fp, pools); err != nil {
		panic(err)
	}
}
