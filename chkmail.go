package chkmail

import "regexp"

// Check Email Expression content with @ and one .
func ChkReg(n string) bool {
	var emailRegex = regexp.MustCompile("^[a-z0-9]+@+[a-z0-9]+\\.[a-z0-9.][a-z0-9]")
	return emailRegex.MatchString(n)
}