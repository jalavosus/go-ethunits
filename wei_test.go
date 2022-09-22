package ethunits_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jalavosus/go-ethunits"
)

func assertBigIntEqual(t *testing.T, want, got *big.Int) bool {
	t.Helper()
	eq := want.Cmp(got) == 0
	return assert.Truef(t, eq, "expected %[1]s to equal %[2]s", got.String(), want.String())
}

func TestToWei_String_Unit(t *testing.T) {
	tests := []testCase[*big.Int, string, ethunits.Unit]{
		{
			name: "unit=ether,amount=" + fromEtherStr,
			args: testArgs[string, ethunits.Unit]{
				amount:   fromEtherStr,
				fromUnit: ethunits.Ether,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=szabo,amount=" + fromSzaboStr,
			args: testArgs[string, ethunits.Unit]{
				amount:   fromSzaboStr,
				fromUnit: ethunits.Szabo,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=gwei,amount=" + fromGweiStr,
			args: testArgs[string, ethunits.Unit]{
				amount:   fromGweiStr,
				fromUnit: ethunits.GWei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=mwei,amount=" + fromMweiStr,
			args: testArgs[string, ethunits.Unit]{
				amount:   fromMweiStr,
				fromUnit: ethunits.MWei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=kwei,amount=" + fromKweiStr,
			args: testArgs[string, ethunits.Unit]{
				amount:   fromKweiStr,
				fromUnit: ethunits.KWei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=wei,amount=" + wantWeiStr,
			args: testArgs[string, ethunits.Unit]{
				amount:   wantWeiStr,
				fromUnit: ethunits.Wei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
	}

	for _, tt := range tests {
		testToWei(t, tt)
	}
}

func TestToWei_String_Uint8(t *testing.T) {
	tests := []testCase[*big.Int, string, uint8]{
		{
			name: "unit=1,amount=" + fromEtherStr,
			args: testArgs[string, uint8]{
				amount:   fromEtherStr,
				fromUnit: 1,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=6,amount=" + fromSzaboStr,
			args: testArgs[string, uint8]{
				amount:   fromSzaboStr,
				fromUnit: 6,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=9,amount=" + fromGweiStr,
			args: testArgs[string, uint8]{
				amount:   fromGweiStr,
				fromUnit: 9,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=12,amount=" + fromMweiStr,
			args: testArgs[string, uint8]{
				amount:   fromMweiStr,
				fromUnit: 12,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=15,amount=" + fromKweiStr,
			args: testArgs[string, uint8]{
				amount:   fromKweiStr,
				fromUnit: 15,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=18,amount=" + wantWeiStr,
			args: testArgs[string, uint8]{
				amount:   wantWeiStr,
				fromUnit: 18,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=17,amount=" + wantWeiStr,
			args: testArgs[string, uint8]{
				amount:   wantWeiStr,
				fromUnit: 17,
			},
			want:   nil,
			wantOk: false,
		},
	}

	for _, tt := range tests {
		testToWei(t, tt)
	}
}

func TestToWei_String_Int(t *testing.T) {
	tests := []testCase[*big.Int, string, int]{
		{
			name: "unit=1,amount=" + fromEtherStr,
			args: testArgs[string, int]{
				amount:   fromEtherStr,
				fromUnit: 1,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=6,amount=" + fromSzaboStr,
			args: testArgs[string, int]{
				amount:   fromSzaboStr,
				fromUnit: 6,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=9,amount=" + fromGweiStr,
			args: testArgs[string, int]{
				amount:   fromGweiStr,
				fromUnit: 9,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=12,amount=" + fromMweiStr,
			args: testArgs[string, int]{
				amount:   fromMweiStr,
				fromUnit: 12,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=15,amount=" + fromKweiStr,
			args: testArgs[string, int]{
				amount:   fromKweiStr,
				fromUnit: 15,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=18,amount=" + wantWeiStr,
			args: testArgs[string, int]{
				amount:   wantWeiStr,
				fromUnit: 18,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=17,amount=" + wantWeiStr,
			args: testArgs[string, int]{
				amount:   wantWeiStr,
				fromUnit: 17,
			},
			want:   nil,
			wantOk: false,
		},
	}

	for _, tt := range tests {
		testToWei(t, tt)
	}
}

func TestToWei_BigInt_Unit(t *testing.T) {
	tests := []testCase[*big.Int, *big.Int, ethunits.Unit]{
		{
			name: "unit=ether,amount=" + fromEtherStr2,
			args: testArgs[*big.Int, ethunits.Unit]{
				amount:   makeBigInt(fromEtherStr2),
				fromUnit: ethunits.Ether,
			},
			want:   makeBigInt(wantWeiStr2),
			wantOk: true,
		},
		{
			name: "unit=szabo,amount=" + fromSzaboStr,
			args: testArgs[*big.Int, ethunits.Unit]{
				amount:   makeBigInt(fromSzaboStr),
				fromUnit: ethunits.Szabo,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=gwei,amount=" + fromGweiStr,
			args: testArgs[*big.Int, ethunits.Unit]{
				amount:   makeBigInt(fromGweiStr),
				fromUnit: ethunits.GWei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=mwei,amount=" + fromMweiStr,
			args: testArgs[*big.Int, ethunits.Unit]{
				amount:   makeBigInt(fromMweiStr),
				fromUnit: ethunits.MWei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=kwei,amount=" + fromKweiStr,
			args: testArgs[*big.Int, ethunits.Unit]{
				amount:   makeBigInt(fromKweiStr),
				fromUnit: ethunits.KWei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=wei,amount=" + wantWeiStr,
			args: testArgs[*big.Int, ethunits.Unit]{
				amount:   makeBigInt(wantWeiStr),
				fromUnit: ethunits.Wei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
	}

	for _, tt := range tests {
		testToWei(t, tt)
	}
}

func TestToWei_BigInt_Uint8(t *testing.T) {
	tests := []testCase[*big.Int, *big.Int, uint8]{
		{
			name: "unit=1,amount=" + fromEtherStr2,
			args: testArgs[*big.Int, uint8]{
				amount:   makeBigInt(fromEtherStr2),
				fromUnit: 1,
			},
			want:   makeBigInt(wantWeiStr2),
			wantOk: true,
		},
		{
			name: "unit=6,amount=" + fromSzaboStr,
			args: testArgs[*big.Int, uint8]{
				amount:   makeBigInt(fromSzaboStr),
				fromUnit: 6,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=9,amount=" + fromGweiStr,
			args: testArgs[*big.Int, uint8]{
				amount:   makeBigInt(fromGweiStr),
				fromUnit: 9,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=12,amount=" + fromMweiStr,
			args: testArgs[*big.Int, uint8]{
				amount:   makeBigInt(fromMweiStr),
				fromUnit: 12,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=15,amount=" + fromKweiStr,
			args: testArgs[*big.Int, uint8]{
				amount:   makeBigInt(fromKweiStr),
				fromUnit: 15,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=18,amount=" + wantWeiStr,
			args: testArgs[*big.Int, uint8]{
				amount:   makeBigInt(wantWeiStr),
				fromUnit: 18,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=17,amount=" + wantWeiStr,
			args: testArgs[*big.Int, uint8]{
				amount:   makeBigInt(wantWeiStr),
				fromUnit: 17,
			},
			want:   nil,
			wantOk: false,
		},
	}

	for _, tt := range tests {
		testToWei(t, tt)
	}
}

func TestToWei_BigInt_Int(t *testing.T) {
	tests := []testCase[*big.Int, *big.Int, int]{
		{
			name: "unit=1,amount=" + fromEtherStr2,
			args: testArgs[*big.Int, int]{
				amount:   makeBigInt(fromEtherStr2),
				fromUnit: 1,
			},
			want:   makeBigInt(wantWeiStr2),
			wantOk: true,
		},
		{
			name: "unit=6,amount=" + fromSzaboStr,
			args: testArgs[*big.Int, int]{
				amount:   makeBigInt(fromSzaboStr),
				fromUnit: 6,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=9,amount=" + fromGweiStr,
			args: testArgs[*big.Int, int]{
				amount:   makeBigInt(fromGweiStr),
				fromUnit: 9,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=12,amount=" + fromMweiStr,
			args: testArgs[*big.Int, int]{
				amount:   makeBigInt(fromMweiStr),
				fromUnit: 12,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=15,amount=" + fromKweiStr,
			args: testArgs[*big.Int, int]{
				amount:   makeBigInt(fromKweiStr),
				fromUnit: 15,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=18,amount=" + wantWeiStr,
			args: testArgs[*big.Int, int]{
				amount:   makeBigInt(wantWeiStr),
				fromUnit: 18,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=17,amount=" + wantWeiStr,
			args: testArgs[*big.Int, int]{
				amount:   makeBigInt(wantWeiStr),
				fromUnit: 17,
			},
			want:   nil,
			wantOk: false,
		},
	}

	for _, tt := range tests {
		testToWei(t, tt)
	}
}

func TestToWei_BigFloat_Unit(t *testing.T) {
	tests := []testCase[*big.Int, *big.Float, ethunits.Unit]{
		{
			name: "unit=ether,amount=" + fromEtherStr,
			args: testArgs[*big.Float, ethunits.Unit]{
				amount:   makeBigFloat(fromEtherStr),
				fromUnit: ethunits.Ether,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=szabo,amount=" + fromSzaboStr,
			args: testArgs[*big.Float, ethunits.Unit]{
				amount:   makeBigFloat(fromSzaboStr),
				fromUnit: ethunits.Szabo,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=gwei,amount=" + fromGweiStr,
			args: testArgs[*big.Float, ethunits.Unit]{
				amount:   makeBigFloat(fromGweiStr),
				fromUnit: ethunits.GWei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=mwei,amount=" + fromMweiStr,
			args: testArgs[*big.Float, ethunits.Unit]{
				amount:   makeBigFloat(fromMweiStr),
				fromUnit: ethunits.MWei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=kwei,amount=" + fromKweiStr,
			args: testArgs[*big.Float, ethunits.Unit]{
				amount:   makeBigFloat(fromKweiStr),
				fromUnit: ethunits.KWei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=wei,amount=" + wantWeiStr,
			args: testArgs[*big.Float, ethunits.Unit]{
				amount:   makeBigFloat(wantWeiStr),
				fromUnit: ethunits.Wei,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
	}

	for _, tt := range tests {
		testToWei(t, tt)
	}
}

func TestToWei_BigFloat_Uint8(t *testing.T) {
	tests := []testCase[*big.Int, *big.Float, uint8]{
		{
			name: "unit=1,amount=" + fromEtherStr,
			args: testArgs[*big.Float, uint8]{
				amount:   makeBigFloat(fromEtherStr),
				fromUnit: 1,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=6,amount=" + fromSzaboStr,
			args: testArgs[*big.Float, uint8]{
				amount:   makeBigFloat(fromSzaboStr),
				fromUnit: 6,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=9,amount=" + fromGweiStr,
			args: testArgs[*big.Float, uint8]{
				amount:   makeBigFloat(fromGweiStr),
				fromUnit: 9,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=12,amount=" + fromMweiStr,
			args: testArgs[*big.Float, uint8]{
				amount:   makeBigFloat(fromMweiStr),
				fromUnit: 12,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=15,amount=" + fromKweiStr,
			args: testArgs[*big.Float, uint8]{
				amount:   makeBigFloat(fromKweiStr),
				fromUnit: 15,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=18,amount=" + wantWeiStr,
			args: testArgs[*big.Float, uint8]{
				amount:   makeBigFloat(wantWeiStr),
				fromUnit: 18,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=17,amount=" + wantWeiStr,
			args: testArgs[*big.Float, uint8]{
				amount:   makeBigFloat(wantWeiStr),
				fromUnit: 17,
			},
			want:   nil,
			wantOk: false,
		},
	}

	for _, tt := range tests {
		testToWei(t, tt)
	}
}

func TestToWei_BigFloat_Int(t *testing.T) {
	tests := []testCase[*big.Int, *big.Float, int]{
		{
			name: "unit=1,amount=" + fromEtherStr,
			args: testArgs[*big.Float, int]{
				amount:   makeBigFloat(fromEtherStr),
				fromUnit: 1,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=6,amount=" + fromSzaboStr,
			args: testArgs[*big.Float, int]{
				amount:   makeBigFloat(fromSzaboStr),
				fromUnit: 6,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=9,amount=" + fromGweiStr,
			args: testArgs[*big.Float, int]{
				amount:   makeBigFloat(fromGweiStr),
				fromUnit: 9,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=12,amount=" + fromMweiStr,
			args: testArgs[*big.Float, int]{
				amount:   makeBigFloat(fromMweiStr),
				fromUnit: 12,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=15,amount=" + fromKweiStr,
			args: testArgs[*big.Float, int]{
				amount:   makeBigFloat(fromKweiStr),
				fromUnit: 15,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=18,amount=" + wantWeiStr,
			args: testArgs[*big.Float, int]{
				amount:   makeBigFloat(wantWeiStr),
				fromUnit: 18,
			},
			want:   makeBigInt(wantWeiStr),
			wantOk: true,
		},
		{
			name: "unit=17,amount=" + wantWeiStr,
			args: testArgs[*big.Float, int]{
				amount:   makeBigFloat(wantWeiStr),
				fromUnit: 17,
			},
			want:   nil,
			wantOk: false,
		},
	}

	for _, tt := range tests {
		testToWei(t, tt)
	}
}

func testToWei[T1 amountConstraint, T2 unitConstraint](t *testing.T, tt testCase[*big.Int, T1, T2]) {
	t.Helper()
	t.Run(tt.name, func(t *testing.T) {
		got, gotOk := ethunits.ToWei(tt.args.amount, tt.args.fromUnit)
		assert.Equal(t, tt.wantOk, gotOk)
		if tt.want != nil {
			assertBigIntEqual(t, tt.want, got)
		}
	})
}