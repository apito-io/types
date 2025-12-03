# YAML Tag Injection for Protobuf Files

## Problem

The protobuf compiler generates Go structs with only `protobuf` and `json` tags. When using `yaml.v3` to unmarshal YAML files into these structs, the unmarshaling fails because yaml.v3 doesn't automatically use json tags as a fallback.

## Solution

We've implemented an automated solution to inject `yaml` tags into the generated protobuf files.

### Steps

1. **Add @inject_tag comments in proto files** (`plugin.proto`)

   - Comments are added to document where yaml tags should go
   - Format: `// @inject_tag: yaml:"field_name"`

2. **Generate protobuf files**

   - Run: `protoc --go_out=. --go-grpc_out=. plugin.proto`

3. **Inject yaml tags**

   - Run: `go run inject_yaml_tags.go`
   - This script reads the generated `.pb.go` files and adds yaml tags based on json tags

4. **Automated with Makefile**
   - Simply run: `make plugin`
   - This runs all steps automatically

### Script: `inject_yaml_tags.go`

The script:

- Reads `protobuff/plugin.pb.go`
- Finds all struct fields with protobuf and json tags
- Adds corresponding yaml tags
- Writes the modified file back

### Config File Structure

Plugin config.yml files should match the protobuf schema exactly. **Use readable string values** for language and placement:

```yaml
id: "plugin-id"
language: "go" # Readable string: "go", "js", "python", "typescript", etc.
title: "Plugin Title"
version: "1.0.0"
enable: true
debug: false
binary_path: "plugin-binary"
handshake_config:
  protocol_version: 1
  magic_cookie_key: "APITO_PLUGIN"
  magic_cookie_value: "apito_plugin_magic_cookie_v1"
env_vars:
  - key: "VAR_NAME"
    value: "var_value"
graphql_schema_config:
  queries:
    - name: "queryName"
      placement: "external" # Readable string: "internal" or "external"
  mutations:
    - name: "mutationName"
      placement: "external"
rest_api_config: # Direct array, no "routes:" wrapper
  - route: "/path"
    placement: "external"
```

### Important Notes

1. **String values**: Use readable strings for `language` and `placement` (not numeric enums)
   - Language: "go", "js", "python", "typescript", "cpp", "java", "ruby", "php", "csharp", "rust", "lua", "dart"
   - Placement: "internal" or "external"
2. **REST API Config**: Direct array under `rest_api_config`, no nested `routes:` key
3. **Field names**: Use json tag names from the protobuf (snake_case)
4. **No wrapper**: Don't wrap config in a `plugin:` key at the root level

### Regenerating Protobuf Files

Whenever you modify `plugin.proto`, run:

```bash
make plugin
```

This will:

1. Generate new .pb.go files
2. Automatically inject yaml tags
3. Ready for YAML unmarshaling

### Testing

You can test YAML unmarshaling with:

```go
import (
	"gopkg.in/yaml.v3"
	"gitlab.com/apito.io/types/protobuff"
)

func TestUnmarshal() {
	data, _ := os.ReadFile("config.yml")
	var config protobuff.PluginDetails
	err := yaml.Unmarshal(data, &config)
	// Should work without errors now!
}
```
