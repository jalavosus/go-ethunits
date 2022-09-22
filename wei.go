package ethunits

import (
	"math/big"
)

func ToWei[T CurrencyAmount, U CurrencyUnit](amount T, fromUnit U) (*big.Int, bool) {
	amt := parseCurrencyAmount(amount)
	unitFrom, ok := parseCurrencyUnit(fromUnit)
	if !ok {
		return nil, ok
	}

	return amountToWei(amt, unitFrom), true
}

func amountToWei(amount *big.Float, fromUnit Unit) *big.Int {
	var (
		exp int
		wei *big.Int
	)

	switch fromUnit {
	case Wei:
		wei, _ = amount.Int(nil)
		return wei
	case Ether:
		exp = 18
	default:
		exp = 18 - unitDecimalsMap[fromUnit]
	}

	powVal := pow10(exp)

	amt := amount.Mul(amount, powVal)
	wei, _ = amt.Int(nil)

	return wei
}