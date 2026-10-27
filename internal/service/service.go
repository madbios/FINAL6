package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoSelect(s string) (string, error) {
	morseFlag := true
	hasMorseChar := false
	if strings.TrimSpace(s) == "" {
		return "", errors.New("пустая строка")
	}
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' {
			morseFlag = false
		}
		if r == '.' || r == '-' {
			hasMorseChar = true
		}
	}
	if morseFlag && hasMorseChar {
		return ConvToText(s)
	}
	return ConvToMorse(s)
}
func ConvToMorse(s string) (string, error) {
	if strings.TrimSpace(s) == "" {
		return "", errors.New("пустая строка")
	}
	for _, r := range s {
		if morse.RuneToMorse(r) == "" {
			return "", fmt.Errorf("символ не поддерживается: %q", r)
		}
	}
	return morse.ToMorse(s), nil
}
func ConvToText(s string) (string, error) {
	if strings.TrimSpace(s) == "" {
		return "", errors.New("пустая строка")
	}
	sSlice := strings.Fields(s)
	for _, sl := range sSlice {
		if morse.MorseToRune(sl) == 0 {
			return "", fmt.Errorf("неизвестный код Морзе: %q ", sl)
		}
	}
	return morse.ToText(s), nil
}
