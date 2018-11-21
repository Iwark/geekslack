package geekslack

import "strings"

func contains(text string, texts []string) bool {
	for _, s := range texts {
		if strings.Contains(text, s) {
			return true
		}
	}
	return false
}
