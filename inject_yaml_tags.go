package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	filePath := "./protobuff/plugin.pb.go"

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	content := string(data)

	// Regex to match struct field tags
	// Pattern: Field name followed by protobuf tag with json tag
	re := regexp.MustCompile(`(\w+)\s+([^\s]+)\s+` + "`" + `(protobuf:"[^"]+"\s+json:"([^"]+)")` + "`")

	// Replace with yaml tag added
	newContent := re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 5 {
			return match
		}

		fieldName := parts[1]
		fieldType := parts[2]
		existingTags := parts[3]
		jsonTag := parts[4]

		// Don't add yaml tag if field is state, unknownFields, or sizeCache
		if fieldName == "state" || fieldName == "unknownFields" || fieldName == "sizeCache" {
			return match
		}

		// Add yaml tag
		newTags := existingTags + ` yaml:"` + jsonTag + `"`
		return fieldName + " " + fieldType + " `" + newTags + "`"
	})

	// Write back
	if err := os.WriteFile(filePath, []byte(newContent), 0644); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	// Count how many tags were added
	oldCount := strings.Count(content, "yaml:")
	newCount := strings.Count(newContent, "yaml:")
	fmt.Printf("✓ Successfully injected yaml tags into %s\n", filePath)
	fmt.Printf("  Added %d yaml tags\n", newCount-oldCount)
}
