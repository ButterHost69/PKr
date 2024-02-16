package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

var (
	KEY_SIZE = 4096
)

func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, KEY_SIZE)
	if err != nil {
		fmt.Println(" ~ Could not create Keys")
		return nil, nil
	}

	return privateKey, &privateKey.PublicKey
}

func ParsePrivateKeyToBytes(pkey *rsa.PrivateKey) []byte {
	pkeyBytes := x509.MarshalPKCS1PrivateKey(pkey)
	privatekey_pem_block := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: pkeyBytes,
		},
	)

	return privatekey_pem_block

}

func ParsePublicKeyToBytes(pbkey *rsa.PublicKey) []byte {
	pbkeyBytes := x509.MarshalPKCS1PublicKey(pbkey)
	publickey_pem_block := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pbkeyBytes,
		},
	)

	return publickey_pem_block
}

func StorePrivateKeyInFile(filepath string, pkey *rsa.PrivateKey) error {
	private_pem_key := ParsePrivateKeyToBytes(pkey)

	if private_pem_key == nil {
		return errors.New("~ Private Key Could Not Be Converted To []Byte")
	}

	return os.WriteFile(filepath, private_pem_key,0666)
}

func StorePublicKeyInFile(filepath string, pbkey *rsa.PublicKey) error {
	public_pem_key := ParsePublicKeyToBytes(pbkey)

	if public_pem_key == nil {
		return errors.New("~ Private Key Could Not Be Converted To []Byte")
	}

	return os.WriteFile(filepath, public_pem_key, 0666)
}
