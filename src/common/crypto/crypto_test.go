package crypto

import (
	"crypto/ecdsa"
	"testing"
)

// A basic test, checks for the crypto constants
func TestConstants(t *testing.T) {
	/*if VerifyPublicKeySize() != true {
		t.Fatal("PublicKeySize does not match libsodium crypto_sign_PUBLICKEYBYTES")
	}

	if VerifySecretKeySize() != true {
		t.Fatal("SecretKeySize does not match libsodium crypto_sign_SECRETKEYBYTES")
	}

	if VerifySignatureSize() != true {
		t.Fatal("SignatureSize does not match libsodium crypto_sign_BYTES")
	}*/

	// Check hash size agasint the libsodium constant for hash size
	if VerifyHashSize() != true {
		t.Fatal("HashSize does not match libsodium crpyto_hash_BYTES")
	}
}
