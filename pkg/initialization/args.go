package initialization

import (
	"errors"
	"strings"
)

// ValidateArgs is in charge of validating command line arguments.
func ValidateArgs(args []string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}

	firstPossibleArg := "init, new, list, list-tags"

	if !strings.Contains(firstPossibleArg, args[0]) {
		return errors.New("unrecognized argumetn")
	}

	return nil
}
