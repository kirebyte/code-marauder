package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kirebyte/code-marauder/settings"
)

func main() {
	// Initialize the application configuration
	cfg := settings.Runtime
	sourceCodeDirs := cfg.CodeListingDirs
	outputFile := cfg.CodeListingFile

	// Create the output directory
	if err := createOutputDir(outputFile); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	// Create the output file
	if err := createOutputFile(outputFile); err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}

	// Walk through the source code directories
	for _, dir := range sourceCodeDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		}

		// Find files and process them
		if err := processDirectory(dir, outputFile, cfg.CodeListingAllowedFiles); err != nil {
			fmt.Println("Error walking through directory:", err)
		}
	}

	fmt.Println("code listing generated at ", outputFile)
}

func createOutputDir(outputFile string) error {
	outputDir := filepath.Dir(outputFile)
	return os.MkdirAll(outputDir, os.ModePerm)
}

func createOutputFile(outputFile string) error {
	return os.WriteFile(outputFile, []byte{}, 0644)
}

func processDirectory(dir, outputFile string, allowedFiles []string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only files
		if !info.IsDir() {
			ext := filepath.Ext(path)
			// Check if the file extension is allowed
			if isAllowedFile(ext, allowedFiles) {
				processFile(path, outputFile)
			}
		}
		return nil
	})
}

func isAllowedFile(ext string, allowedFiles []string) bool {
	for _, allowedExt := range allowedFiles {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

// Function to process each file
func processFile(filePath, outputFile string) {
	// Read the file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Open the output file in append mode
	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}
	defer f.Close()

	// Append the file content to the output file
	fileContent := fmt.Sprintf("==== File: %s ====\n%s", filePath, string(data)) // "==== File: <filePath> ====\n<fileContent>"
	_, err = f.WriteString(fileContent)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
	}

	// Append a separator to the output file
	separator := "==== End of File ====\n\n"
	_, err = f.WriteString(separator)
	if err != nil {
		fmt.Println("Error writing separator to output file:", err)
	}
}
