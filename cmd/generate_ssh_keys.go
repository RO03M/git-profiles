package cmd

import (
	"crypto/ed25519"
	"encoding/pem"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

func GenerateKey(email string, password string) ([]byte, []byte) {
	pubKey, privKey, _ := ed25519.GenerateKey(nil)

	publicKey, _ := ssh.NewPublicKey(pubKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	publicKeyWithEmail := fmt.Sprintf("%s %s", strings.ReplaceAll(string(authorizedKey), "\n", ""), email)

	fmt.Println(publicKeyWithEmail)

	var pemPrivateKey *pem.Block

	if password == "" {
		pemPrivateKey, _ = ssh.MarshalPrivateKey(privKey, "aes256-ctr")
	} else {
		pemPrivateKey, _ = ssh.MarshalPrivateKeyWithPassphrase(privKey, "aes256-ctr", []byte(password))
	}

	return []byte(publicKeyWithEmail), pem.EncodeToMemory(pemPrivateKey)
}
