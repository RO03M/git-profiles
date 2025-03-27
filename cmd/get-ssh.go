package cmd

import (
	"crypto/ed25519"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	"github.com/mikesmitty/edkey"
	"github.com/orochaa/go-clack/prompts"
	"golang.org/x/crypto/ssh"
)

func SanitizeSshFileName(filename string) string {
	filename = strings.Replace(filename, "@", "a", -1)
	filename = strings.Replace(filename, ".", "", -1)

	return filename
}

func CreateNewSshFile(email string) {
	pubKey, privKey, _ := ed25519.GenerateKey(nil)

	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(privKey),
	}

	publicKey, _ := ssh.NewPublicKey(pubKey)
	privateKey := pem.EncodeToMemory(pemKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	fmt.Println(string(privateKey), string(authorizedKey))

	homeDir, _ := os.UserHomeDir()

	path := fmt.Sprintf("%s/.ssh/%s", homeDir, SanitizeSshFileName(email))

	os.WriteFile(path, []byte(string(authorizedKey)+email), 0644)
}

func GetSshPath(email string) string {
	shouldGenerateSsh, _ := prompts.Confirm(prompts.ConfirmParams{
		Message: "Do you wish to generate a ssh key pair?",
	})

	if shouldGenerateSsh {

	}

	return ""
}
