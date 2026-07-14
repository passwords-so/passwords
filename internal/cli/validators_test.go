package cli

import "testing"

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		ruleset Ruleset
		wantErr string
	}{
		{
			name:  "accepts input without rules",
			input: "anything",
		},
		{
			name:    "rejects empty required input",
			ruleset: Ruleset{Required: true},
			wantErr: "input is required",
		},
		{
			name:    "accepts empty optional input",
			ruleset: Ruleset{MinLength: 3, CharSet: ALPHA_CHARSET},
		},
		{
			name:    "rejects input below minimum length",
			input:   "ab",
			ruleset: Ruleset{MinLength: 3},
			wantErr: "input must be at least 3 characters long",
		},
		{
			name:    "accepts input at minimum length",
			input:   "abc",
			ruleset: Ruleset{MinLength: 3},
		},
		{
			name:    "rejects input above maximum length",
			input:   "abcd",
			ruleset: Ruleset{MaxLength: 3},
			wantErr: "input must be at most 3 characters long",
		},
		{
			name:    "accepts input at maximum length",
			input:   "abc",
			ruleset: Ruleset{MaxLength: 3},
		},
		{
			name:    "counts unicode characters",
			input:   "åß",
			ruleset: Ruleset{MinLength: 2, MaxLength: 2},
		},
		{
			name:    "accepts characters in charset",
			input:   "abc123",
			ruleset: Ruleset{CharSet: ALPHA_CHARSET + NUMERIC_CHARSET},
		},
		{
			name:    "rejects character outside charset",
			input:   "abc!",
			ruleset: Ruleset{CharSet: ALPHA_CHARSET},
			wantErr: "input contains invalid character '!'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateInput(test.input, test.ruleset)
			if test.wantErr == "" {
				if err != nil {
					t.Fatalf("ValidateInput() error = %v, want nil", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("ValidateInput() error = nil, want %q", test.wantErr)
			}
			if err.Error() != test.wantErr {
				t.Fatalf("ValidateInput() error = %q, want %q", err, test.wantErr)
			}
		})
	}
}
