package split

import "strings"

// split package with a single split function.

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s, seq string) (result []string) {
	i := strings.Index(s, seq)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(seq):]
		i = strings.Index(s, seq)
	}
	result = append(result, s)
	return 
}
