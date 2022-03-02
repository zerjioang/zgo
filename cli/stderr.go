package cli

import "os"

// ExitWithError generates an error message to stderr and exists the application abruptly
func ExitWithError(err error) {
	if err == nil {
		return
	}
	_, _ = os.Stderr.WriteString(err.Error())
	os.Exit(1)
}
