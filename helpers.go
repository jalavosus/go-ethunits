package ethunits

import (
	"math/big"
)

func parseCurrencyAmount[T CurrencyAmount](amount T) (amt *big.Float) {
	switch x := (any)(amount).(type) {
	case *big.Float:
		amt = x
	case *big.Int:
		amt = new(big.Float).SetInt(x)
	case string:
		amt, _ = new(big.Float).SetString(x)
	}

	return
}

func parseCurrencyUnit[T CurrencyUnit](fromUnit T) (unitFrom Unit, ok bool) {
	switch x := (any)(fromUnit).(type) {
	case Unit:
		unitFrom = x
		ok = true
	case uint8:
		unitFrom, ok = UnitFromDecimals(x)
	case int:
		unitFrom, ok = UnitFromDecimals(x)
	}

	return
}