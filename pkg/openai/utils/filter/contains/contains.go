package contains

func StrValue(val string, list []string) bool {
	for i := range list {
		if list[i] == val {
			return true
		}
	}
	return false
}
