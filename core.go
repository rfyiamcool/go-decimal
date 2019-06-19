package decimal

func IsSorted(vals []*Decimal) bool {
	for idx, _ := range vals {
		if idx+1 > len(vals)-1 {
			break
		}

		if vals[idx].GTcmp(vals[idx+1]) {
			return false
		}
	}

	return true
}

func IsReverseSorted(vals []*Decimal) bool {
	for idx, _ := range vals {
		if idx+1 > len(vals)-1 {
			break
		}

		if vals[idx].LTcmp(vals[idx+1]) {
			return false
		}
	}

	return true
}
