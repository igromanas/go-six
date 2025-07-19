package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(s string) string {
	if isMorse(s) {
		return morse.ToText(s)
	} else {
		return morse.ToMorse(s)
	}
}

func isMorse(s string) bool {
	for _, ch := range s {
		if ch != ' ' && ch != '.' && ch != '-' {
			return false
		}
	}

	return true
}
