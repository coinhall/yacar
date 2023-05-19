package enums

type Chain string

const (
	Osmosis      Chain = "osmosis"
	Juno         Chain = "juno"
	Kujira       Chain = "kujira"
	Terra        Chain = "terra"
	TerraClassic Chain = "terraclassic"
)

func (c Chain) Name() string {
	return string(c)
}

func getAllChains() []Chain {
	return []Chain{
		Osmosis,
		Juno,
		Kujira,
		Terra,
		TerraClassic,
	}
}

func GetAllChainNames() []string {
	var names []string
	for _, chain := range getAllChains() {
		names = append(names, chain.Name())
	}
	return names
}
