package aes

import (
	"fmt"

	openssl "github.com/Luzifer/go-openssl/v4"
)

func decrypt() {
	// encrypt: echo "test@2023" | openssl aes-256-cbc -pbkdf2 -md sha256 -a
	// decrypt: echo "U2FsdGVkX1/oep17EAD8EUjAQAU8jUG21jOI9Wg7YF8=" | openssl enc -aes-256-cbc -pbkdf2 -md sha256 -a -d
	opensslEncrypted := "U2FsdGVkX1/oep17EAD8EUjAQAU8jUG21jOI9Wg7YF8="
	passphrase := "123456"

	o := openssl.New()

	dec, err := o.DecryptBytes(passphrase, []byte(opensslEncrypted), openssl.PBKDF2SHA256)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}

	fmt.Printf("Decrypted text: %s\n", string(dec))
}
