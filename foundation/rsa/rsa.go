package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func GetPublicKey(path string) (*rsa.PublicKey, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(buf)

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, nil
}

func GetPrivateKey(path string) (*rsa.PrivateKey, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(buf)

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func RSA_Sign(hash []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
	if err != nil {
		return nil, err
	}
	return sign, nil
}

func RSA_Verify(hash []byte, sign []byte, publicKey *rsa.PublicKey) error {
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash, sign)
}

func main() {
	publicKey, err := GetPublicKey("public.pem")
	if err != nil {
		fmt.Println("GetPublicKey error: ", err)
		return
	}
	privateKey, err := GetPrivateKey("private.pem")
	if err != nil {
		fmt.Println("GetPrivateKey error: ", err)
		return
	}

	plainText := []byte("hello world")

	encryptText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		fmt.Println("rsa.EncryptPKCS1v15 error: ", err)
		return
	}
	fmt.Println("encrypt: ", encryptText)

	decryptText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptText)
	if err != nil {
		fmt.Println("rsa.DecryptPKCS1v15 error: ", err)
		return
	}
	fmt.Println("decrypt: ", string(decryptText))

	sha256Hash := sha256.Sum256(plainText)
	fmt.Println("hash: ", sha256Hash)

	signedHash, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, sha256Hash[:])
	if err != nil {
		fmt.Println("rsa.SignPKCS1v15 error: ", err)
		return
	}
	fmt.Println("sign: ", signedHash)

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, sha256Hash[:], signedHash)
	if err != nil {
		fmt.Println("rsa.VerifyPKCS1v15 failed: ", err)
		return
	}
	fmt.Println("verify sign success")

}
