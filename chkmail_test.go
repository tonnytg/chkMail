package chkmail

import (
	"fmt"
	"testing"
)

// Test validations it's working
func TestChkReg(t *testing.T) {
	mail := "test@domain.com"
	m := ChkReg(mail)
	fmt.Println(m)
	if ! m {
		t.Errorf("Email not pass")
	}
}
