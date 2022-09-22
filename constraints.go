package ethunits

import (
	"math/big"
)

type (
	CurrencyAmount interface{ *big.Int | *big.Float | string }
	CurrencyUnit   interface{ Unit | uint8 | int }
	DecimalValue   interface{ uint8 | int }
)