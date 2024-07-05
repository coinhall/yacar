package sorter

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/coinhall/yacar/internal/unmarshaler"
	"github.com/coinhall/yacar/internal/walker"
	"github.com/coinhall/yacarsdk/v2"
)

func Start(filePaths []string) {
	sortYacarJSONs(filePaths)
	log.Println("Sorted JSONs successfully...")
}

func sortYacarJSONs(filePaths []string) {
	for _, fp := range filePaths {
		fp = filepath.ToSlash(fp)
		switch walker.MustParse(walker.TrimJsonSuffixFromFullPath(fp)) {
		case walker.Account:
			handleAccount(fp)
		case walker.Asset:
			handleAsset(fp)
		case walker.Binary:
			handleBinary(fp)
		case walker.Contract:
			handleContract(fp)
		case walker.Entity:
			handleEntity(fp)
		case walker.Pool:
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
	if err := writeFile(fp, accounts); err != nil {
		panic(err)
	}
}

func handleAsset(fp string) {
	assets, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Asset, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedAssetOrder(assets))
	if err := writeFile(fp, assets); err != nil {
		panic(err)
	}
}

func handleBinary(fp string) {
	binaries, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Binary, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedBinaryOrder(binaries))
	if err := writeFile(fp, binaries); err != nil {
		panic(err)
	}
}

func handleContract(fp string) {
	contracts, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Contract, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedContractOrder(contracts))
	if err := writeFile(fp, contracts); err != nil {
		panic(err)
	}
}

func handleEntity(fp string) {
	entities, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Entity, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedEntityOrder(entities))
	if err := writeFile(fp, entities); err != nil {
		panic(err)
	}
}

func handlePool(fp string) {
	pools, err := unmarshaler.UnmarshalInto(fp, make([]yacarsdk.Pool, 0))
	if err != nil {
		panic(err)
	}
	sort.Stable(yacarsdk.ByEnforcedPoolOrder(pools))
	if err := writeFile(fp, pools); err != nil {
		panic(err)
	}
}

func writeFile[T any](path string, data []T) error {
	var sb strings.Builder
	sbEnc := json.NewEncoder(&sb)
	sbEnc.SetEscapeHTML(false)
	sbEnc.SetIndent("", "  ")
	if err := sbEnc.Encode(data); err != nil {
		panic(err)
	}

	parentDir := filepath.Dir(path)
	if err := os.MkdirAll(parentDir, 0o755); err != nil {
		return err
	}

	// If file exists, overwrite the contents completely
	if err := os.WriteFile(path, []byte(sb.String()), 0o644); err != nil {
		return err
	}

	return nil
}
