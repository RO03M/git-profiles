package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
	"log"

	"github.com/orochaa/go-clack/prompts"
)

func ExecAddRoutine(defaultConfig config.Config) {
	name, err := prompts.Text(prompts.TextParams{
		Message:     "Git account name",
		Placeholder: "Mona Lisa",
	})

	if err != nil {
		log.Fatalf("Failed to get the git account name %v\n", err)
	}

	email, err := prompts.Text(prompts.TextParams{
		Message:     "Git account email",
		Placeholder: "yourgitaccount@email.com",
	})

	if err != nil {
		log.Fatalf("Failed to get the git account email %v\n", err)
	}

	sshPath := GetSshPath(email)

	// GetSshKey()
	fmt.Println(sshPath)

	if err != nil {
		log.Fatalf("Failed to get the absolute ssh path %v\n", err)
	}

	profileName, err := prompts.Text(prompts.TextParams{
		Message:     "What should this git profile be called?",
		Required:    false,
		Placeholder: fmt.Sprintf("Profile for %s", email),
	})

	if err != nil {
		log.Fatalf("Failed to get the profile name %v\n", err)
	}

	var profile = config.Profile{
		ProfileName:     profileName,
		Name:            name,
		Email:           email,
		AbsoluteSshPath: absoluteSshPath,
	}

	defaultConfig.Profiles = append(defaultConfig.Profiles, profile)

	defaultConfig.Save()
}
