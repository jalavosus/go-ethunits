package ethunits

import (
	"math/big"
)

func ToEther[T CurrencyAmount, U CurrencyUnit](amount T, fromUnit U) (*big.Float, bool) {
	amt := parseCurrencyAmount(amount)
	unitFrom, ok := parseCurrencyUnit(fromUnit)
	if !ok {
		return nil, ok
	}

	return amountToEther(amt, unitFrom), true
}

func amountToEther(amount *big.Float, fromUnit Unit) *big.Float {
	var wei *big.Int

	switch fromUnit {
	case Ether:
		return amount
	default:
		wei = amountToWei(amount, fromUnit)
	}

	f := new(big.Float)
	fWei := new(big.Float).SetInt(wei)
	amt := f.Quo(fWei, pow10(18))

	return amt
}