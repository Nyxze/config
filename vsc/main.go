package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
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
