package gitconfig

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
}

func GetConfigPath() string {
	userHomeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalf("Failed to get the user homedir %v\n", err)
		panic(err)
	}

	return fmt.Sprintf("%s/%s", userHomeDir, ".gitconfig")
}
