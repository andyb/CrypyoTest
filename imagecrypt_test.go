package imagecrypt

import (
	"log"
	"testing"
)

func TestEncyptImage(t *testing.T) {
	msg := "This is the message to be encrypted"
	key := "1234567890123456"
	encrypted := make([]byte, len(msg))
	err := Encrypt(encrypted, []byte(msg), []byte(key))

	if err != nil {
		panic(err)
	}

	
	log.Printf("Encrypted value is %s", encrypted)

	decrypted := make([]byte, len(msg))
	err = Decrypt(decrypted, encrypted, []byte(key))
	if err != nil {
		panic(err)
	}
	
	log.Printf("Decrypted value is %s", decrypted)
}
