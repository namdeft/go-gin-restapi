package validation

import "regexp"

func ValidatePassword(password string) bool {
	return len(password) >= 7 &&
		contains(password, "[0-9]") &&
		contains(password, "[A-Z]") &&
		contains(password, `[!@#$%^&*()_+{}|:"<>?~]`)
}

func contains(s, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}
