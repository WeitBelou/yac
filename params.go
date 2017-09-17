package yac

import (
	"fmt"
	"strings"
)

type Params map[string]string

const ParamsContextKey = "_params"

// Extracts params determined by pattern from given path
func extractParams(pattern, path string) (Params, error) {
	if !matchPattern(pattern, path) {
		return nil, fmt.Errorf("'%s' doesn't match pattern '%s'", path, pattern)
	}

	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	params := Params{}

	for i, part := range patternParts {
		if part != pathParts[i] {
			name, err := extractParamName(part)
			if err != nil {
				return nil, fmt.Errorf("invalid pattern format: '%s'", pattern)
			}

			params[name] = pathParts[i]
		}
	}

	return params, nil
}

// Checks if given path matches pattern
func matchPattern(pattern, path string) bool {
	if pattern == path {
		return true
	}

	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return false
	}

	for i, part := range patternParts {
		if part != pathParts[i] {
			if !isParamName(part) {
				return false
			}
		}
	}

	return true
}

// Helper to determine is given string can be param name.
func isParamName(s string) bool {
	return strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}")
}

// Helper to extract param name from string.
func extractParamName(s string) (string, error) {
	if !isParamName(s) {
		return "", fmt.Errorf("'%s' is not a param name", s)
	}
	if s == "{}" {
		return "", fmt.Errorf("'%s' has empty name part", s)
	}

	return s[1: len(s)-1], nil
}
