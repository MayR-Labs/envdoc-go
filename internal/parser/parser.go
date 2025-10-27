package parser

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// EnvVar represents an environment variable
type EnvVar struct {
	Key     string
	Value   string
	Comment string
}

// ParseEnvFile parses a .env file and returns a list of environment variables
func ParseEnvFile(filename string) ([]EnvVar, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	var envVars []EnvVar
	scanner := bufio.NewScanner(file)
	var currentComment string

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// Handle comments
		if strings.HasPrefix(trimmed, "#") {
			currentComment = trimmed
			continue
		}

		// Skip empty lines
		if trimmed == "" {
			currentComment = ""
			continue
		}

		// Parse key=value
		parts := strings.SplitN(trimmed, "=", 2)
		if len(parts) == 2 {
			envVar := EnvVar{
				Key:     strings.TrimSpace(parts[0]),
				Value:   strings.TrimSpace(parts[1]),
				Comment: currentComment,
			}
			envVars = append(envVars, envVar)
			currentComment = ""
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return envVars, nil
}

// WriteEnvFile writes environment variables to a file
func WriteEnvFile(filename string, envVars []EnvVar) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	writer := bufio.NewWriter(file)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			return
		}
	}(writer)

	for _, envVar := range envVars {
		if envVar.Comment != "" {
			_, err := fmt.Fprintln(writer, envVar.Comment)

			if err != nil {
				return err
			}
		}

		_, err := fmt.Fprintf(writer, "%s=%s\n", envVar.Key, envVar.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetEnvKeys returns a list of unique keys from environment variables
func GetEnvKeys(envVars []EnvVar) []string {
	keys := make([]string, len(envVars))
	for i, envVar := range envVars {
		keys[i] = envVar.Key
	}
	return keys
}

// ArrangeByPrefix groups environment variables by their prefix
func ArrangeByPrefix(envVars []EnvVar) []EnvVar {
	if len(envVars) == 0 {
		return envVars
	}

	// Sort by key
	sort.Slice(envVars, func(i, j int) bool {
		return envVars[i].Key < envVars[j].Key
	})

	return envVars
}

// FindDuplicates finds duplicate keys in environment variables
func FindDuplicates(envVars []EnvVar) []string {
	keyCount := make(map[string]int)
	for _, envVar := range envVars {
		keyCount[envVar.Key]++
	}

	var duplicates []string
	for key, count := range keyCount {
		if count > 1 {
			duplicates = append(duplicates, key)
		}
	}

	sort.Strings(duplicates)
	return duplicates
}

// FindMissingKeys finds keys that are in one set but not another
func FindMissingKeys(source, target []string) []string {
	targetMap := make(map[string]bool)
	for _, key := range target {
		targetMap[key] = true
	}

	var missing []string
	for _, key := range source {
		if !targetMap[key] {
			missing = append(missing, key)
		}
	}

	sort.Strings(missing)
	return missing
}
