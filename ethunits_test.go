package ethunits_test

import (
	"math/big"

	"github.com/jalavosus/go-ethunits"
)

const (
	wantWeiStr    string = "342500000000000000000"
	wantWeiStr2   string = "342000000000000000000"
	wantEtherStr  string = "342.5"
	fromEtherStr  string = "342.5"
	fromEtherStr2 string = "342"
	fromSzaboStr  string = "342500000"
	fromGweiStr   string = "342500000000"
	fromMweiStr   string = "342500000000000"
	fromKweiStr   string = "342500000000000000"
)

type (
	amountConstraint interface{ *big.Int | *big.Float | string }
	unitConstraint   interface{ ethunits.Unit | uint8 | int }
	wantConstraint   interface{ *big.Int | *big.Float }
)

type testArgs[T amountConstraint, U unitConstraint] struct {
	amount   T
	fromUnit U
}

type testCase[T1 wantConstraint, T2 amountConstraint, T3 unitConstraint] struct {
	name   string
	args   testArgs[T2, T3]
	want   T1
	wantOk bool
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