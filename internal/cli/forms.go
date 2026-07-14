package cli

import (
	"errors"

	"charm.land/huh/v2"
)

func InitVaultForm() (string, string, error) {
	var name, password, confirmPassword string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Vault name").
				Value(&name).
				Validate(required),
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
		return "", "", err
	}

	return name, password, nil
}

func validatePassword(value string) error {
	if err := required(value); err != nil {
		return err
	}

	if len(value) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}

func required(value string) error {
	if value == "" {
		return errors.New("this field is required")
	}

	return nil
}
