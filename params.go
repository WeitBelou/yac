package yac

import "regexp"

type Params map[string]string

var paramRegexp = regexp.MustCompile(`{([[:word:]]+?)}`)

// Replaces all quasi patterns with regexp named groups
func patternToRegexp(pattern string) string {
	return paramRegexp.ReplaceAllString(pattern, `(?P<$1>[[:alnum:]]+?)`)
}
