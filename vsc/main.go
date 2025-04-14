package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	filesToCopy = []string{"keybindings.json", "settings.json"}
)

func main() {
	// Define flags
	var fileFlag bool
	var extFlag bool
	flag.BoolVar(&fileFlag, "f", false, "Update files to copy")
	flag.BoolVar(&extFlag, "ext", false, "Update extension file")

	flag.Parse()

	if fileFlag {
		updateFiles()
	}
	if extFlag {
		updateExtension()
	}
}

func updateExtension() {
	// Read JSON content from file
	content, err := os.ReadFile("ext.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse JSON into a slice of maps
	var data []map[string]interface{}
	if err := json.Unmarshal(content, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Extract extension IDs
	extensionIDs := extractExtensionIDs(data)

	// Write the extension IDs to a file
	if err := writeToFile("extension.txt", extensionIDs); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
func updateFiles() {
	home := os.Getenv("HOME")
	home = strings.ReplaceAll(home, "\\", "/")
	path := home + "/AppData/Roaming/Code/User/"
	for i := 0; i < len(filesToCopy); i++ {
		src := path + filesToCopy[i]
		dest := "Settings/" + filesToCopy[i]
		copyFile(src, dest)
	}
}
func copyFile(srcPath, destPath string) error {
	// Open the source file
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()

	// Create the destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destFile.Close()

	// Copy the file content
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}

	return nil
}

// extractExtensionIDs pulls out the VSCode extension IDs from the JSON
func extractExtensionIDs(data []map[string]interface{}) []string {
	var result []string

	for _, entry := range data {
		if identifier, ok := entry["identifier"].(map[string]interface{}); ok {
			if id, ok := identifier["id"].(string); ok {
				result = append(result, id)
			}
		}
	}
	return result
}

// writeToFile writes each line from a string slice to a file, one per line
func writeToFile(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}
