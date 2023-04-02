package validator

func CheckLengthWordAndTranslate(word, translate string) bool {
	if len(word) < 1 || len(translate) < 1 {
		return false
	}
	return true
}
