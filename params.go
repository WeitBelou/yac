package yac

import "strings"

type Params map[string]string

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
			if !strings.HasPrefix(part, "{") || !strings.HasSuffix(part, "}") {
				return false
			}
		}
	}

	return true
}
