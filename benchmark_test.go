package chkmail

import (
	"testing"
)

// BenchmarkChkReg mede a performance da função ChkReg.
func BenchmarkChkReg(b *testing.B) {
	validEmail := "test@example.com"
	invalidEmail := "invalid-email@"

	for i := 0; i < b.N; i++ {
		ChkReg(validEmail)
		ChkReg(invalidEmail)
	}
}

func BenchmarkChkRegMultiple(b *testing.B) {
	emails := []string{
		"test@example.com",
		"another.test@domain.co.uk",
		"user123@example.org",
		"invalid-email@",
		"username@domain..com",
	}

	for i := 0; i < b.N; i++ {
		for _, email := range emails {
			ChkReg(email)
		}
	}
}
