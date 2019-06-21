# go-decimal

golang easy decimal lib

`inf module copy from gopkg.in/inf.v0`

## usage:

```
package decimal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSub(t *testing.T) {
	var (
		balance, _    = NewDecimalFromString("8.001")
		addBalance, _ = NewDecimalFromString("1.001")
		subBalance, _ = NewDecimalFromString("3.001")

		addres, _ = NewDecimalFromString("9.002")
		subres, _ = NewDecimalFromString("6.001")
	)

	balance.Add(balance, addBalance)
	assert.Equal(t, addres, balance)

	balance.Sub(balance, subBalance)
	assert.Equal(t, balance, subres)
}

func TestMulQuo(t *testing.T) {
	var (
		balance       = NewDecimalFromInt64(88)
		const2Decimal = NewDecimalFromInt64(2)
		res = NewDecimalFromInt64(44)

		minPre        = NewDecimalFromInt64(1)
	)

	assert.Equal(t,
		res,
		new(Decimal).QuoRoundFloor(balance, const2Decimal, minPre),
	)

	assert.Equal(t,
		balance,
		new(Decimal).Mul(const2Decimal, res),
	)
}

func TestCmp(t *testing.T) {
	var (
		a, _ = NewDecimalFromString("8.001")
		b, _ = NewDecimalFromString("1.001")
	)

	assert.Equal(t, true, a.GTcmp(b))
	assert.Equal(t, false, a.LTcmp(b))
}

func TestPrecision(t *testing.T) {
	var (
		prec, _   = NewDecimalFromString("0.001")
		source, _ = NewDecimalFromString("123.1234567890")
		dst, _    = NewDecimalFromString("123.123")
	)

	assert.Equal(
		t,
		dst,
		new(Decimal).RoundFloor(source, prec),
	)
}
```