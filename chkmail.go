package chkmail

import (
	"regexp"
	"strings"
)

// ChkReg will check regex for Email Expression content with @ and one .
func ChkReg(n string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(n) {
		return false
	}
	// check for invalid domain patterns
	var invalidDomainPatterns = regexp.MustCompile(`(^-|-$|\.{2,})`)
	domain := n[strings.Index(n, "@")+1:]
	return !invalidDomainPatterns.MatchString(domain)
}
