package cmd

import (
	"crypto/ed25519"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
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
	// privKey = ed25519.Sign(privKey, []byte("teste"))
	// pemKey := &pem.Block{
	// 	Type:  "OPENSSH PRIVATE KEY",
	// 	Bytes: edkey.MarshalED25519PrivateKey(privKey),
	// }

	publicKey, _ := ssh.NewPublicKey(pubKey)
	// privateKey := pem.EncodeToMemory(pemKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	homeDir, _ := os.UserHomeDir()

	filename := fmt.Sprintf("%s%s", SanitizeSshFileName(email), uuid.New())

	err := os.MkdirAll(fmt.Sprintf("%s/.ssh/gitprofiles/", homeDir), os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	publicKeyPath := fmt.Sprintf("%s/.ssh/gitprofiles/%s.pub", homeDir, filename)
	privateKeyPath := fmt.Sprintf("%s/.ssh/gitprofiles/%s", homeDir, filename)

	publicKeyContent := fmt.Sprintf("%s %s", strings.ReplaceAll(string(authorizedKey), "\n", ""), email)

	// signer, _ := ssh.NewSignerFromKey(privKey)
	encryptedPEM, err := ssh.MarshalPrivateKeyWithPassphrase(privKey, "aes256-ctr", []byte("vsauce"))

	if err != nil {
		panic(err)
	}

	fmt.Println(encryptedPEM)

	os.WriteFile(publicKeyPath, []byte(publicKeyContent), 0600)
	os.WriteFile(privateKeyPath, []byte(pem.EncodeToMemory(encryptedPEM)), 0600)

	return publicKeyPath, privateKeyPath
}

func GetSshPath(email string) (string, string, bool) {
	shouldGenerateSsh, _ := prompts.Confirm(prompts.ConfirmParams{
		Message: "Do you wish to generate a ssh key pair?",
	})

	if shouldGenerateSsh {
		publicPath, privatePath := CreateNewSshFile(email)
		prompts.Info(fmt.Sprintf("Created the public ssh file at\n%s", publicPath))

		return publicPath, privatePath, shouldGenerateSsh
	}

	absoluteSshPath, _ := prompts.Path(prompts.PathParams{
		Message:  "Write the path to the ssh credential",
		Required: true,
	})

	return absoluteSshPath, "", shouldGenerateSsh
}
