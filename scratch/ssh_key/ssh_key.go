package sshkey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

func generateKeyPair() (string, string, string) {
	// Generate a new RSA key pair with a key size of 2048 bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Encode the private key in PEM format
	privateKeyBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	// Generate the public key from the private key
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	// Calculate the fingerprint of the public key
	// This is equivalent to running `ssh-keygen -lf <public key file>` in the command line
	fingerprint := ssh.FingerprintLegacyMD5(publicKey)

	return string(privateKeyBytes), string(ssh.MarshalAuthorizedKey(publicKey)), fingerprint
}
