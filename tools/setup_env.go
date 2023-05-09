package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open .env file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Printf("Invalid line in .env file: %s\n", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		err := os.Setenv(key, value)
		if err != nil {
			fmt.Printf("Failed to set env variables: %s\n", line)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading .env file: %v\n", err)
	}
}
