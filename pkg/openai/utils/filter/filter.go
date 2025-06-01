package filter

import "github.com/kylerqws/chatbot/pkg/openai/utils/filter/contains"

func CheckStrValue(val string, list []string) bool {
	return len(list) > 0 && !contains.StrValue(val, list)
}

func CheckDateValue(date, after, before int64) bool {
	return date > 0 && ((after > 0 && date <= after) || (before > 0 && date >= before))
}
