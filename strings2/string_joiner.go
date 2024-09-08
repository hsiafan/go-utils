package strings2

import "strings"

// Joiner is joiner setting for join/split string
type Joiner struct {
	Prefix    string
	Suffix    string
	Separator string
}

// Join joins string
func (j *Joiner) Join(ss []string) string {
	var sb strings.Builder

	totalLen := 0
	totalLen += len(j.Prefix)
	for i, s := range ss {
		if i > 0 {
			totalLen += len(j.Separator)
		}
		totalLen += len(s)
	}
	totalLen += len(j.Suffix)

	sb.Grow(totalLen)

	sb.WriteString(j.Prefix)
	for i, s := range ss {
		if i > 0 {
			sb.WriteString(j.Separator)
		}
		sb.WriteString(s)
	}
	sb.WriteString(j.Suffix)

	return sb.String()
}

// Split splits string into items
func (j *Joiner) Split(s string) []string {
	if len(j.Prefix) > 0 && strings.HasPrefix(s, j.Prefix) {
		s = s[len(j.Prefix):]
	}
	if len(j.Suffix) > 0 && strings.HasSuffix(s, j.Suffix) {
		s = s[:len(s)-len(j.Suffix)]
	}
	return strings.Split(s, j.Separator)
}
