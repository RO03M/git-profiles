package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
	"git-profiles/cmd/utils"
	"log"
	"os"

	"github.com/orochaa/go-clack/prompts"
	"gopkg.in/ini.v1"
)

func SwitchProfileRoutine(configFile config.Config) {
	home, _ := os.UserHomeDir()

	cfg, err := ini.Load(home + "/.gitconfig")
	if err != nil {
		panic(err)
	}

	// fmt.Println(cfg.Section("user").Key("email"))
	profiles := utils.Map(configFile.Profiles, func(item config.Profile, key int) *prompts.SelectOption[config.Profile] {
		return &prompts.SelectOption[config.Profile]{
			Label: item.ProfileName,
			Value: item,
			Hint:  fmt.Sprintf("%s %s", item.Name, item.Email),
		}
	})

	if len(profiles) == 0 {
		prompts.Info("No profiles found, create one using the add command")
		return
	}

	targetProfile, err := prompts.Select(prompts.SelectParams[config.Profile]{
		Options: profiles,
	})

	if err != nil {
		log.Fatalf("Failed to select a profile %v\n", err)
		return
	}

	fmt.Println(targetProfile)
	cfg.Section("user").Key("email").SetValue(targetProfile.Email)
	cfg.Section("user").Key("name").SetValue(targetProfile.Name)
	cfg.Section("core").Key("sshCommand").SetValue(fmt.Sprintf("ssh -i %s", targetProfile.AbsoluteSshPath))

	cfg.SaveTo(home + "/.gitconfig")
}
