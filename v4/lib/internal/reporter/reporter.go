package reporter

import (
	"fmt"
	"os"
)

// ReportError centralizes error reporting for unexpected errors.
func ReportError(err error, context string) {
	if err == nil {
		return
	}
	// Centralized error reporting: log to stderr with context.
	fmt.Fprintf(os.Stderr, "ERROR [%s]: %v\n", context, err)
}
