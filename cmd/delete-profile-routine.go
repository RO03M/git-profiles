package cmd

import (
	"fmt"
	"git-profiles/cmd/config"
	"git-profiles/cmd/utils"
	"os"

	"github.com/orochaa/go-clack/prompts"
)

func DeleteProfileRoutine(configFile config.Config) {
	var deleteParam string
	var err error

	if len(os.Args) >= 3 {
		deleteParam = os.Args[2]
	} else {
		profiles := utils.Map(configFile.Profiles, func(item config.Profile, key int) *prompts.SelectOption[string] {
			return &prompts.SelectOption[string]{
				Label: item.ProfileName,
				Value: item.Id.String(),
				Hint:  fmt.Sprintf("%s %s", item.Name, item.Email),
			}
		})

		profiles = append(profiles, &prompts.SelectOption[string]{
			Label: "Cancel",
			Value: "-1",
		})

		deleteParam, err = prompts.Select(prompts.SelectParams[string]{
			Message: "Choose a profile to be deleted",
			Options: profiles,
		})
	}

	if deleteParam == "-1" || prompts.IsCancel(err) {
		fmt.Println("Canceled deletion operation")
		return
	}

	profileToBeDeleted := utils.Find(configFile.Profiles, func(profile config.Profile, _ int) bool {
		return profile.Id.String() == deleteParam || profile.ProfileName == deleteParam || profile.Name == deleteParam || profile.Email == deleteParam
	})

	deleteConfirmation, err := prompts.Confirm(prompts.ConfirmParams{
		Message: fmt.Sprintf("Are you sure you want to delete the profile %s (email: %s, name: %s)", profileToBeDeleted.ProfileName, profileToBeDeleted.Email, profileToBeDeleted.Name),
	})

	if err != nil || prompts.IsCancel(err) || !deleteConfirmation {
		prompts.Error("Canceled the delete operation!")

		return
	}

	profiles := utils.Filter(configFile.Profiles, func(profile config.Profile, _ int) bool {
		return profile.Id != profileToBeDeleted.Id
	})

	configFile.Profiles = profiles
	configFile.Save()
}
