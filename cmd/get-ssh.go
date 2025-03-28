package cmd

import (
	"crypto/ed25519"
	"encoding/pem"
	"fmt"
	"log"
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

func CreateNewSshFile(email string) (string, string) {
	pubKey, privKey, _ := ed25519.GenerateKey(nil)

	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(privKey),
	}

	publicKey, _ := ssh.NewPublicKey(pubKey)
	privateKey := pem.EncodeToMemory(pemKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	homeDir, _ := os.UserHomeDir()

	filename := SanitizeSshFileName(email)

	err := os.MkdirAll(fmt.Sprintf("%s/.ssh/gitprofiles/", homeDir), os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	publicKeyPath := fmt.Sprintf("%s/.ssh/gitprofiles/%s.pub", homeDir, filename)
	privateKeyPath := fmt.Sprintf("%s/.ssh/gitprofiles/%s", homeDir, filename)

	publicKeyContent := fmt.Sprintf("%s %s", strings.ReplaceAll(string(authorizedKey), "\n", " "), email)

	os.WriteFile(publicKeyPath, []byte(publicKeyContent), 0644)
	os.WriteFile(privateKeyPath, []byte(privateKey), 0644)

	return publicKeyPath, privateKeyPath
}

func GetSshPath(email string) (string, bool) {
	shouldGenerateSsh, _ := prompts.Confirm(prompts.ConfirmParams{
		Message: "Do you wish to generate a ssh key pair?",
	})

	if shouldGenerateSsh {
		publicPath, _ := CreateNewSshFile(email)
		prompts.Info(fmt.Sprintf("Created the public ssh file at\n%s", publicPath))

		return publicPath, shouldGenerateSsh
	}

	absoluteSshPath, _ := prompts.Path(prompts.PathParams{
		Message:  "Write the path to the ssh credential",
		Required: true,
	})

	return absoluteSshPath, shouldGenerateSsh
}
