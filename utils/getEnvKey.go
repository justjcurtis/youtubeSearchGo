/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"errors"
	"os"
	"strings"
)

func GetEnvKey(key string) (string, error) {
	env, err := os.ReadFile(".env")
	if err != nil {
		return err.Error(), err
	}
	lines := strings.Split(string(env), "\n")
	entries := make(map[string]string)
	for _, line := range lines {
		entry := strings.Split(line, "=")
		if len(entry) != 2 {
			continue
		}
		entries[entry[0]] = entry[1]
	}
	result, ok := entries[key]
	if !ok {
		return "", errors.New("Key not found")
	}
	return result, nil
}
