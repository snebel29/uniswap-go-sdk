package uniswap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrency_NewCurrency(t *testing.T) {
	c := NewCurrency(18, "DAI", "Dai")
	assert.EqualValues(t, c.Decimals, 18)
	assert.Equal(t, c.Symbol, "DAI")
	assert.Equal(t, c.Name, "Dai")
}
