package yacar

type YacarFile string

const (
	Account  YacarFile = "account"
	Asset    YacarFile = "asset"
	Binary   YacarFile = "binary"
	Contract YacarFile = "contract"
	Entity   YacarFile = "entity"
	Pool     YacarFile = "pool"
)

const FileSuffix = ".json"

func MustParse(file string) YacarFile {
	switch file {
	case "account":
		return Account
	case "asset":
		return Asset
	case "binary":
		return Binary
	case "contract":
		return Contract
	case "entity":
		return Entity
	case "pool":
		return Pool
	default:
		panic("unhandled file case: " + file)
	}
}

func GetAllFiles() []YacarFile {
	return []YacarFile{
		Account,
		Asset,
		Binary,
		Contract,
		Entity,
		Pool,
	}
}

func GetAllFilesWithExt() []string {
	fileEnums := GetAllFiles()
	s := make([]string, 0, len(fileEnums))
	for _, f := range fileEnums {
		s = append(s, string(f)+".json")
	}
	return s
}
