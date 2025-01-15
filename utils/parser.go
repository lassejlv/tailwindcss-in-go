package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func Parser(config Config) {
	extensions := config.Extensions
	var matches []string
	var allClasses []string

	// Loop through each extension and append matches
	for _, ext := range extensions {
		if !strings.Contains(ext, "html") {
			log.Warn().Msgf("Skipping extension %s (not supported at the moment)", ext)
			continue
		}
		m, err := filepath.Glob(ext)
		if err != nil {
			log.Error().Msgf("Error getting files: %s", err)
			continue
		}
		if len(m) == 0 {
			log.Warn().Msgf("No files found for extension %s", ext)
			continue
		}
		matches = append(matches, m...)
	}

	if len(matches) == 0 {
		log.Error().Msg("No files found")
		return
	}

	for _, match := range matches {
		content, err := os.ReadFile(match)
		if err != nil {
			log.Error().Msgf("Error reading file %s: %s", match, err)
			continue
		}

		// Find all classes in "class="
		parts := strings.Split(string(content), "class=\"")
		for _, part := range parts[1:] { // Skip first part (before first class=")
			// Find the closing quote
			if idx := strings.Index(part, "\""); idx != -1 {
				classString := part[:idx]
				// Split the class string by spaces to get individual classes
				classes := strings.Fields(classString) // This handles multiple spaces/tabs/newlines
				allClasses = append(allClasses, classes...)
			}
		}
	}

	// Filter out duplicates
	uniqueClasses := make(map[string]bool)
	for _, class := range allClasses {
		uniqueClasses[class] = true
	}

	// Convert map to slice
	uniqueClassesSlice := make([]string, 0, len(uniqueClasses))
	for class := range uniqueClasses {
		uniqueClassesSlice = append(uniqueClassesSlice, class)
	}

	// Write to file
	file, err := os.Create("classes.txt")
	if err != nil {
		log.Error().Msgf("Error creating file: %s", err)
		return
	}
	defer file.Close()

	for _, class := range uniqueClassesSlice {
		_, err := file.WriteString(class + "\n")
		if err != nil {
			log.Error().Msgf("Error writing to file: %s", err)
			return
		}
	}

	log.Info().Msgf("Found %d unique classes in %d files", len(uniqueClassesSlice), len(matches))
}
