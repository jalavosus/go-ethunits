package ethunits

//go:generate stringer -type Unit

type Unit uint

const (
	// Wei is the base currency unit of ethereum-based blockchains,
	// also equivalent to 10^-18 Ether.
	Wei Unit = iota
	// KWei (KiloWei) is equivalent to 10^-15 Ether.
	KWei
	// MWei (MegaWei) is equivalent to 10^-12 Ether.
	MWei
	// GWei (GigaWei) is equivalent to 10^-9 Ether.
	GWei
	// Szabo is equivalent to 10^-6 Ether.
	// It is used commonly by the USDC and USDT token contracts as their
	// decimal value, for some reason.
	Szabo
	// Finney is equivalent to 10^-3 Ether.
	Finney
	// Ether is commonly used as the "user-facing" representation of
	// amounts of tokens on ethereum-based blockchains.
	// It is equal to 10^18 Wei.
	Ether
	// Unknown represents an unknown Unit.
	Unknown
)

func (u Unit) Int() int {
	return int(u)
}

func (u Unit) Uint() uint {
	return uint(u)
}

// var unitWeiExponentMap = map[Unit]int{
// 	Ether:  18,
// 	Finney: 15,
// 	Szabo:  12,
// 	GWei:   9,
// 	MWei:   6,
// 	KWei:   3,
// 	Wei:    1,
// }

var unitDecimalsMap = map[Unit]int{
	Ether:  1,
	Finney: 3,
	Szabo:  6,
	GWei:   9,
	MWei:   12,
	KWei:   15,
	Wei:    18,
}

func UnitFromDecimals[T DecimalValue](decimals T) (Unit, bool) {
	var (
		u  Unit
		ok bool
	)

	for k, v := range unitDecimalsMap {
		if v == int(decimals) {
			u = k
			ok = true
			break
		}
	}

	return u, ok
}