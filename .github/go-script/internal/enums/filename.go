package enums

type FileName string

const (
	Account  FileName = "account"
	Asset    FileName = "asset"
	Binary   FileName = "binary"
	Contract FileName = "contract"
	Entity   FileName = "entity"
	Pool     FileName = "pool"
)

func (f FileName) Name() string {
	return string(f)
}

func getAllFiles() []FileName {
	return []FileName{
		Account,
		Asset,
		Binary,
		Contract,
		Entity,
		Pool,
	}
}

func GetAllFileNames() []string {
	var names []string
	for _, file := range getAllFiles() {
		names = append(names, file.Name())
	}
	return names
}
