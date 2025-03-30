package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
	"log"
	"os"

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

	sshPath, generatedSshKey := GetSshPath(email)

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
		AbsoluteSshPath: sshPath,
	}

	defaultConfig.Profiles = append(defaultConfig.Profiles, profile)

	defaultConfig.Save()

	if generatedSshKey {
		pubKey, _ := os.ReadFile(sshPath)
		prompts.Note(string(pubKey), prompts.NoteOptions{
			Title: "Paste the following key into your git provider",
		})

		prompts.Select(prompts.SelectParams[string]{
			Message: "Did you paste it into the git provider?",
			Options: []*prompts.SelectOption[string]{
				{
					Label: "Yes",
				},
				{
					Label: "I'll do it later, trust me :)",
				},
			},
		})
	}

	prompts.Info("New Profile created! Use git-profiles su and choose it to use")

}
