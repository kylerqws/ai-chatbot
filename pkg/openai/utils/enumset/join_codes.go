package enumset

import (
	"sort"
	"strings"
)

// JoinCodes extracts and joins sorted map keys using the given separator.
func JoinCodes[M ~map[string]V, V any](m M, sep string) string {
	codes := make([]string, 0, len(m))
	for k := range m {
		codes = append(codes, k)
	}
	sort.Strings(codes)
	return strings.Join(codes, sep)
}
