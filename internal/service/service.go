package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
)

// AutoConvert определяет формат строки и конвертирует её.
// Если передан код Морзе — возвращает текст, иначе — код Морзе.
func AutoConvert(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return ""
	}

	if isMorseCode(input) {
		return morse.ToText(input)
	}
	return morse.ToMorse(input)
}

// isMorseCode проверяет, состоит ли строка из символов Морзе
func isMorseCode(s string) bool {
	morseChars := 0
	totalChars := 0
	for _, ch := range s {
		if ch == '.' || ch == '-' || ch == '/' || ch == ' ' {
			morseChars++
		}
		if ch != '\n' && ch != '\r' {
			totalChars++
		}
	}
	if totalChars == 0 {
		return false
	}
	// если более 80% символов — символы Морзе
	return float64(morseChars)/float64(totalChars) > 0.8
}
