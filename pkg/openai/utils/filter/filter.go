package filter

import "slices"

func MatchStrValue(val string, list []string) bool {
	if len(list) == 0 {
		return true
	}
	return slices.Contains(list, val)
}

func MatchDateValue(date, after, before int64) bool {
	if date <= 0 {
		return true
	}

	if after > 0 && date <= after {
		return false
	}
	if before > 0 && date >= before {
		return false
	}

	return true
}
