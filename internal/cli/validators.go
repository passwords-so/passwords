package cli

import "fmt"

type Ruleset struct {
	Required  bool
	MinLength int
	MaxLength int
	CharSet   string
}

const ALPHA_CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NUMERIC_CHARSET = "0123456789"
const SPECIAL_CHARSET = "!@#$%^&*()-_=+[]{}|;:',.<>?/`~"

var VaultPasswordRuleset = Ruleset{
	Required:  true,
	MinLength: 8,
	MaxLength: 64,
	CharSet:   ALPHA_CHARSET + NUMERIC_CHARSET + SPECIAL_CHARSET,
}

var VaultNameRuleset = Ruleset{
	Required:  true,
	MinLength: 1,
	MaxLength: 32,
	CharSet:   ALPHA_CHARSET + NUMERIC_CHARSET + "-_",
}

func ValidateInput(input string, ruleset Ruleset) error {
	if input == "" {
		if ruleset.Required {
			return fmt.Errorf("input is required")
		}
		return nil
	}

	length := len([]rune(input))
	if ruleset.MinLength > 0 && length < ruleset.MinLength {
		return fmt.Errorf("input must be at least %d characters long", ruleset.MinLength)
	}
	if ruleset.MaxLength > 0 && length > ruleset.MaxLength {
		return fmt.Errorf("input must be at most %d characters long", ruleset.MaxLength)
	}

	if ruleset.CharSet != "" {
		for _, char := range input {
			valid := false
			for _, allowed := range ruleset.CharSet {
				if char == allowed {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("input contains invalid character %q", char)
			}
		}
	}

	return nil

}
