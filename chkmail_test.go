package chkmail

import (
	"testing"
)

// TestChkReg check the regex for email expression
func TestChkReg(t *testing.T) {
	validEmails := []string{
		"test@domain.com",
		"another.test@domain.co.uk",
		"user123@example.org",
		"email.with+symbol@example.com",
		"email_with-dash@sub.example.com",
	}

	invalidEmails := []string{
		"plainaddress",
		"@missingusername.com",
		"username@.com",
		"username@com",
		"username@com.",
		"username@-domain.com",
		"username@domain..com",
	}

	for _, email := range validEmails {
		if !ChkReg(email) {
			t.Errorf("Expected valid, but got invalid: %s", email)
		}
	}

	for _, email := range invalidEmails {
		if ChkReg(email) {
			t.Errorf("Expected invalid, but got valid: %s", email)
		}
	}
}
