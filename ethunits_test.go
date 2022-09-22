package ethunits_test

import (
	"math/big"

	"github.com/jalavosus/go-ethunits"
)

type (
	amountConstraint interface{ *big.Int | *big.Float | string }
	unitConstraint   interface{ ethunits.Unit | uint8 | int }
)

type testArgs[T amountConstraint, U unitConstraint] struct {
	amount   T
	fromUnit U
}

func makeBigInt(amt string) *big.Int {
	b := new(big.Int)
	biAmt, _ := b.SetString(amt, 10)
	return biAmt
}

func makeBigFloat(amt string) *big.Float {
	f := new(big.Float)
	bfAmt, _ := f.SetString(amt)
	return bfAmt
}