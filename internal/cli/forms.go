package cli

import (
	"errors"
	"strings"

	"charm.land/huh/v2"
)

// InitVaultForm prompts the user to enter a password and confirm it
func InitVaultForm() (password string, err error) {
	var confirmPassword string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Password").
				EchoMode(huh.EchoModePassword).
				Value(&password).
				Validate(validatePassword),
			huh.NewInput().
				Title("Confirm password").
				EchoMode(huh.EchoModePassword).
				Value(&confirmPassword).
				Validate(func(value string) error {
					if err := validatePassword(value); err != nil {
						return err
					}
					if value != password {
						return errors.New("passwords do not match")
					}
					return nil
				}),
		),
	)

	if err := form.Run(); err != nil {
		return "", err
	}

	return password, nil
}

// a wrapper around cli.ValidateInput to satisy huh
func validatePassword(value string) error {
	if err := ValidateInput(value, VaultPasswordRuleset); err != nil {
		return errors.New(strings.ReplaceAll(err.Error(), "input", "password"))
	}
	return nil
}
