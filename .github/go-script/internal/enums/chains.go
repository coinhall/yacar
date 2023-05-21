package enums

type Chain string

const (
	Juno         Chain = "juno"
	Kujira       Chain = "kujira"
	Osmosis      Chain = "osmosis"
	Terra        Chain = "terra"
	TerraClassic Chain = "terraclassic"
)

func (c Chain) Name() string {
	return string(c)
}

func getAllChains() []Chain {
	return []Chain{
		Juno,
		Kujira,
		Osmosis,
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
