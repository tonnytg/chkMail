package chkmail

import "regexp"

// ChkReg will check regex for Email Expression content with @ and one .
func ChkReg(n string) bool {
	var emailRegex = regexp.MustCompile(`(?!(^[.-]|^.*?[.]{2,}|[.-]$))[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	return emailRegex.MatchString(n)
}
