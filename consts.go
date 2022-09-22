package ethunits

import (
	"math"
	"math/big"
)

var (
	pow18   = math.Pow10(18)
	bfPow18 = big.NewFloat(pow18)
)

var (
	biPow18 *big.Int
)

func init() {
	biPow18, _ = bfPow18.Int(nil)
}

func pow10(exp int) *big.Float {
	return big.NewFloat(math.Pow10(exp))
}