package logging

import (
	"os"
	"regexp"
	"strings"
)

// IsTraceEnabled receives the a short name value (usually the app name) and the fully qualified
// package identifier (i.e. `github.com/dfuse-io/logging/subpackage`) and determines if tracing should
// be enabled for such values.
//
// To take the decision, this function inspects the `TRACE` environment variable. If the variable
// is **not** set, tracing is disabled. If the value is either `true`, `*` or `.*`, then tracing
// is enabled.
//
// For other values, we split the `TRACE` value using `,` separator. For each split element, if
// the element matches directly the short name, tracing is enabled and if the element as a Regexp
// object matches (partially, not fully) the `packageID`, tracing is enabled.
//
// In all other cases, tracing is disabled.
func IsTraceEnabled(shortName string, packageID string) bool {
	trace := os.Getenv("TRACE")
	if trace == "" {
		return false
	}

	if trace == "true" || trace == "TRUE" || trace == "*" || trace == ".*" {
		return true
	}

	for _, part := range strings.Split(trace, ",") {
		if part == shortName {
			return true
		}

		if regexp.MustCompile(part).MatchString(packageID) {
			return true
		}
	}

	return false
}
