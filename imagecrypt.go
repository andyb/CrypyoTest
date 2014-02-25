package imagecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

//1. Encrypt an image
//2. Decrypt an image
//3. Encrypt some text
//4. Decrypt some text

func Encrypt(dst, src, key []byte) error {
	log.Println("Encrypt started")
	iv := []byte(key)[:aes.BlockSize]
	aesBlockEncrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(dst, src)
	log.Println("Encrypt ended")
	return nil
}

func Decrypt(dst, src, key []byte) error {
	iv := []byte(key)[:aes.BlockSize]
	aesBlockDecrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(dst, src)
	return nil
}


