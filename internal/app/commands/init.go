package commands

import (
	"context"
	"fmt"

	"github.com/novembersoftware/passwords/internal/app"
	"github.com/novembersoftware/passwords/internal/utils"
)

func RunInitCmd(ctx context.Context, backend app.Backend, args []string) error {
	name := args[0]
	if err := utils.ValidateInput(name, utils.VaultNameRuleset); err != nil {
		return fmt.Errorf("invalid vault name: %w", err)
	}

	password, err := app.InitVaultForm()
	if err != nil {
		return err
	}
	if err = backend.Create(ctx, name, password); err != nil {
		return err
	}
	return nil
}
