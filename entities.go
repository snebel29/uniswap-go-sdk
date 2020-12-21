package uniswap

// ETHER represents ether.
var ETHER = &Currency{Decimals: 18, Symbol: "ETH", Name: "Ether"}

// Currency is any fungible financial instrument on Ethereum, including Ether and all ERC20 tokens.
type Currency struct {
	// TODO: use bigint?
	Decimals uint8
	Symbol   string
	Name     string
}

// NewCurrency return a new currency.
// TODO: Is using uint8 type to represent decimals correct?
func NewCurrency(decimals uint8, symbol, name string) *Currency {
	return &Currency{
		Decimals: decimals,
		Symbol:   symbol,
		Name:     name,
	}
}
