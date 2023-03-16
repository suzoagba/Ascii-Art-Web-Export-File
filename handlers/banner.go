package handlers

import (
	"os"
)

func ReadBanner(path string) ([]string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data []string
	var line string
	for _, char := range raw {
		if char == '\n' {
			if line == "" {
				continue
			}
			data = append(data, line)
			line = ""
			continue
		}
		line += string(char)
	}

	data = append(data, line)

	return data, nil
}