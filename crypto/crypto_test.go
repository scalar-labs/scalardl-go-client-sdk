package crypto

import (
	"testing"
)

func Test_NewEcSha256Signer_WithIncorrectKey_ShouldGetError(t *testing.T) {
	if _, err := NewEcSha256Signer([]byte("not a key")); err == nil {
		t.Errorf("should get an error")
	}
}

func Test_NewEcSha256Signer_WithCorrectKey_ShouldGetInstance(t *testing.T) {
	var key string = "-----BEGIN EC PRIVATE KEY-----\n" +
		"MHcCAQEEICcJGMEw3dyXUGFu/5a36HqY0ynZi9gLUfKgYWMYgr/IoAoGCCqGSM49\n" +
		"AwEHoUQDQgAEBGuhqumyh7BVNqcNKAQQipDGooUpURve2dO66pQCgjtSfu7lJV20\n" +
		"XYWdrgo0Y3eXEhvK0lsURO9N0nrPiQWT4A==\n-----END EC PRIVATE KEY-----\n"

	if _, err := NewEcSha256Signer([]byte(key)); err != nil {
		t.Errorf("should get an correct signer instance")
	}
}

func Test_EcSha256Signer_WithCorrectKey_ShouldSignCorrectly(t *testing.T) {
	var key string = "-----BEGIN EC PRIVATE KEY-----\n" +
		"MHcCAQEEICcJGMEw3dyXUGFu/5a36HqY0ynZi9gLUfKgYWMYgr/IoAoGCCqGSM49\n" +
		"AwEHoUQDQgAEBGuhqumyh7BVNqcNKAQQipDGooUpURve2dO66pQCgjtSfu7lJV20\n" +
		"XYWdrgo0Y3eXEhvK0lsURO9N0nrPiQWT4A==\n-----END EC PRIVATE KEY-----\n"
	var s EcSha256Signer
	var err error
	var signed []byte

	if s, err = NewEcSha256Signer([]byte(key)); err != nil {
		t.Errorf("should get an correct signer instance")
	}

	if signed, err = s.Sign([]byte("hello world!")); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(signed) < 2 || len(signed) < (int(signed[1])-2) {
		t.Errorf("signature is in an incorrect length")
	}

	if int(signed[0]) != 48 {
		t.Errorf("signature has an incorrect first-byte")
	}
}
