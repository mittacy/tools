package encryption

import "testing"

func TestEncryptionBySalt(t *testing.T) {
	password := "password"
	pw, salt := Encryption(password)
	if pw != EncryptionBySalt(password, salt) {
		t.Errorf("the %s should be %s encrypted with salt(%s)", password, pw, salt)
	}
}
