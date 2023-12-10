package passwd

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

var pubKey, _ = os.ReadFile("files/rsa/rsa.pub")

func parsePubKeyBytes(pub_key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub_key)
	if block == nil {
		return nil, errors.New("block nil")
	}
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

func GenPasswordMD5(passwd string) (string, error) {
	return fmt.Sprintf("%x", md5.Sum([]byte(passwd))), nil
}

func GenPasswordRSA(passwd string) (string, error) {
	pk, err := parsePubKeyBytes(pubKey)
	if err != nil {
		panic(err)
	}
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		pk,
		[]byte(passwd),
		nil)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", encryptedBytes), nil
}
