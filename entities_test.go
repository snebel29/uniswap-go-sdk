package uniswap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrency(t *testing.T) {
	c := NewCurrency(18, "DAI", "Dai")
	assert.Equal(t, c.Decimals, 18)
	assert.Equal(t, c.Symbol, "DAI")
	assert.Equal(t, c.Name, "Dai")
}
