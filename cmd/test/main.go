package main

import (
	"fmt"
	"os"

	"gitlab.com/apito.io/types/protobuff"
	"gopkg.in/yaml.v3"
)

func main() {
	// Read the config file
	data, err := os.ReadFile("/Users/diablo/go/src/gitlab.com/apito.io/udbhabon-plugins/hc-suchok-plugin/config.yml")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// First unmarshal to a map to see the raw data
	var rawMap map[string]interface{}
	if err := yaml.Unmarshal(data, &rawMap); err != nil {
		fmt.Printf("Error unmarshaling to map: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Raw YAML keys: %v\n\n", getKeys(rawMap))

	// Try to unmarshal
	var config protobuff.PluginDetails
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Success! Unmarshaled config:\n")
	fmt.Printf("  ID: %s\n", config.Id)
	fmt.Printf("  Title: %s\n", config.Title)
	fmt.Printf("  Version: %s\n", config.Version)
	fmt.Printf("  Enable: %v\n", config.Enable)
	fmt.Printf("  Debug: %v\n", config.Debug)
	fmt.Printf("  Language: %v\n", config.Language)
	fmt.Printf("  Binary Path: %s\n", config.BinaryPath)
	if config.HandshakeConfig != nil {
		fmt.Printf("  Handshake Config:\n")
		fmt.Printf("    Protocol Version: %d\n", config.HandshakeConfig.ProtocolVersion)
		fmt.Printf("    Magic Cookie Key: %s\n", config.HandshakeConfig.MagicCookieKey)
		fmt.Printf("    Magic Cookie Value: %s\n", config.HandshakeConfig.MagicCookieValue)
	}
	if config.GraphqlSchemaConfig != nil {
		fmt.Printf("  GraphQL Schema Config:\n")
		fmt.Printf("    Queries: %d\n", len(config.GraphqlSchemaConfig.Queries))
		fmt.Printf("    Mutations: %d\n", len(config.GraphqlSchemaConfig.Mutations))
	}
	if config.RestApiConfig != nil {
		fmt.Printf("  REST API Config:\n")
		fmt.Printf("    Routes: %d\n", len(config.RestApiConfig))
		for i, route := range config.RestApiConfig {
			fmt.Printf("      [%d] %s (%s)\n", i+1, route.Route, route.Placement)
		}
	}
	fmt.Printf("  Env Vars: %d\n", len(config.EnvVars))
}

func getKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
