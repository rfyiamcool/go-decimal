package decimal

import (
	inf "gopkg.in/inf.v0"
)

type Decimal struct {
	inf.Dec
}

var Zero = new(Decimal).SetUnscaled(0)
var One = new(Decimal).SetUnscaled(1)

func DecimalComparator(a, b interface{}) int {
	an := a.(*Decimal)
	bn := b.(*Decimal)
	return an.Cmp(bn)
}

func NewDecimalFromString(s string) (*Decimal, bool) {
	nd := new(Decimal)
	_, ok := nd.Dec.SetString(s)
	if ok == false {
		return nil, false
	} else {
		return nd, true
	}
}

func NewDecimalFromInt64(n int64) *Decimal {
	nd := new(Decimal)
	nd.SetUnscaled(n)
	return nd
}

func (d *Decimal) Set(nd *Decimal) *Decimal {
	d.Dec.Set(&nd.Dec)
	return d
}

func (d *Decimal) SetUnscaled(n int64) *Decimal {
	d.Dec.SetUnscaled(n)
	return d
}

func (d1 *Decimal) Abs(d2 *Decimal) *Decimal {
	nd := new(Decimal)
	nd.Dec = *nd.Dec.Abs(&d2.Dec)
	return nd
}

func (d1 *Decimal) Cmp(d2 *Decimal) int {
	return d1.Dec.Cmp(&d2.Dec)
}

func (d1 *Decimal) GTcmp(d2 *Decimal) bool {
	if d1.Dec.Cmp(&d2.Dec) > 0 {
		return true
	}

	return false
}

func (d1 *Decimal) GEcmp(d2 *Decimal) bool {
	if d1.Dec.Cmp(&d2.Dec) >= 0 {
		return true
	}

	return false
}

func (d1 *Decimal) LEcmp(d2 *Decimal) bool {
	if d1.Dec.Cmp(&d2.Dec) <= 0 {
		return true
	}

	return false
}

func (d1 *Decimal) LTcmp(d2 *Decimal) bool {
	if d1.Dec.Cmp(&d2.Dec) < 0 {
		return true
	}

	return false
}

func (d1 *Decimal) Equal(d2 *Decimal) bool {
	if d1.Dec.Cmp(&d2.Dec) == 0 {
		return true
	}

	return false
}

func (d1 *Decimal) NotEqual(d2 *Decimal) bool {
	if d1.Dec.Cmp(&d2.Dec) != 0 {
		return true
	}

	return false
}

func (d1 *Decimal) Sub(d2 *Decimal, d3 *Decimal) *Decimal {
	d1.Dec.Sub(&d2.Dec, &d3.Dec)
	return d1
}

func (d1 *Decimal) Add(d2 *Decimal, d3 *Decimal) *Decimal {
	d1.Dec.Add(&d2.Dec, &d3.Dec)
	return d1
}

func (d1 *Decimal) Mul(d2 *Decimal, d3 *Decimal) *Decimal {
	d1.Dec.Mul(&d2.Dec, &d3.Dec)
	return d1
}

func (d1 *Decimal) MulCeil(d2 *Decimal, d3 *Decimal, base *Decimal) *Decimal {
	d1.Dec.Mul(&d2.Dec, &d3.Dec)
	d1.RoundCeil(d1, base)
	return d1
}

func (d1 *Decimal) MulFloor(d2 *Decimal, d3 *Decimal, base *Decimal) *Decimal {
	d1.Dec.Mul(&d2.Dec, &d3.Dec)
	d1.RoundFloor(d1, base)
	return d1
}

func (d1 *Decimal) CanDivExact(d2 *Decimal) bool {
	if n := new(Decimal).QuoExact(&d1.Dec, &d2.Dec); n == nil {
		return false
	} else if n.Cmp(&One.Dec) < 0 {
		return false
	}

	return true
}

func (d1 *Decimal) QuoRoundFloor(x, y, min *Decimal) *Decimal {
	d1.Dec.QuoRound(&x.Dec, &y.Dec, min.Scale(), inf.RoundFloor)
	return d1
}

func (d1 *Decimal) RoundFloor(x, base *Decimal) *Decimal {
	d1.Dec.Round(&x.Dec, base.Scale(), inf.RoundFloor)
	return d1
}

func (d1 *Decimal) RoundCeil(x, base *Decimal) *Decimal {
	d1.Dec.Round(&x.Dec, base.Scale(), inf.RoundCeil)
	return d1
}

func (d *Decimal) IsZero() bool {
	return d.Dec.Cmp(&Zero.Dec) == 0
}

func (d *Decimal) String() string {
	return d.Dec.String()
}
