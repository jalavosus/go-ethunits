package ethunits

import (
	"math/big"
)

func ToWei[T CurrencyAmount, U CurrencyUnit](amount T, fromUnit U) (*big.Int, bool) {
	var (
		amt      = new(big.Float)
		unitFrom Unit
		ok       bool
	)

	switch x := (any)(amount).(type) {
	case *big.Float:
		amt = x
	case *big.Int:
		amt = amt.SetInt(x)
	case string:
		amt, _ = amt.SetString(x)
	}

	switch x := (any)(fromUnit).(type) {
	case Unit:
		unitFrom = x
		ok = true
	case uint8:
		unitFrom, ok = UnitFromDecimals(x)
	case int:
		unitFrom, ok = UnitFromDecimals(x)
	}

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