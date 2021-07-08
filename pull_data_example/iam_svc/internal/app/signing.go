package app

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type SigningKeyPair struct {
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
}

func GenerateECDSAKeys() SigningKeyPair {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return SigningKeyPair{
		PublicKey:  &privKey.PublicKey,
		PrivateKey: privKey,
	}
}
