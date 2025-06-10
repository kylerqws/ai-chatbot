package filter

import "slices"

// MatchStrValue returns true if val is present in list,
// or if the list is empty (interpreted as "no filtering").
// Returns false if val is nil.
func MatchStrValue(val *string, list []string) bool {
	if len(list) == 0 {
		return true
	}
	if val == nil {
		return false
	}
	return slices.Contains(list, *val)
}

// MatchDateValue returns true if the date is either unset (nil or non-positive),
// or falls strictly after `after` and strictly before `before`, when provided.
func MatchDateValue(date, after, before *int64) bool {
	if date == nil || *date <= 0 {
		return true
	}
	if after != nil && *after > 0 && *date <= *after {
		return false
	}
	if before != nil && *before > 0 && *date >= *before {
		return false
	}
	return true
}
