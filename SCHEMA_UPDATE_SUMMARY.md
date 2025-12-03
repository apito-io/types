# Plugin Schema Update Summary

## Changes Made

### 1. **Language Field: Enum → String**

**Before:**
```protobuf
enum PluginLanguage {
  PLUGIN_LANGUAGE_GO = 0;
  PLUGIN_LANGUAGE_JS = 1;
  // ...
}

message PluginDetails {
  PluginLanguage language = 18;
}
```

**After:**
```protobuf
message PluginDetails {
  string language = 18; // "go", "js", "python", etc.
}
```

**Config YAML:**
```yaml
# Before: language: 0
# After:
language: "go"
```

### 2. **Placement Field: Enum → String**

**Before:**
```protobuf
enum PlacementType {
  PLACEMENT_TYPE_INTERNAL = 0;
  PLACEMENT_TYPE_EXTERNAL = 1;
}

message GraphQLSchemaConfigItem {
  string name = 1;
  PlacementType placement = 2;
}
```

**After:**
```protobuf
message GraphQLSchemaConfigItem {
  string name = 1;
  string placement = 2; // "internal" or "external"
}

message RESTApiConfigItem {
  string route = 1;
  string placement = 2; // "internal" or "external"
}
```

**Config YAML:**
```yaml
# Before: placement: 1
# After:
placement: "external"
```

### 3. **REST API Config: Nested → Direct Array**

**Before:**
```protobuf
message RESTApiConfig {
  repeated RESTApiConfigItem routes = 1;
}

message PluginDetails {
  RESTApiConfig rest_api_config = 23;
}
```

**After:**
```protobuf
message PluginDetails {
  repeated RESTApiConfigItem rest_api_config = 23;
}
```

**Config YAML:**
```yaml
# Before:
rest_api_config:
  routes:
    - route: "/path"
      placement: 1

# After:
rest_api_config:
  - route: "/path"
    placement: "external"
```

## Benefits

1. ✅ **More Readable**: Config files are now human-friendly
2. ✅ **Easier to Edit**: No need to remember numeric enum values
3. ✅ **Self-Documenting**: "external" is clearer than "1"
4. ✅ **Simpler Structure**: Direct arrays instead of nested wrappers
5. ✅ **Source of Truth**: config.yml structure matches exactly what's in the proto

## Supported Values

### Language
- `"go"` or `"golang"`
- `"js"` or `"javascript"`
- `"typescript"` or `"ts"`
- `"python"` or `"py"`
- `"cpp"` or `"c++"`
- `"java"`
- `"ruby"` or `"rb"`
- `"php"`
- `"csharp"` or `"c#"`
- `"rust"`
- `"lua"`
- `"dart"`

### Placement
- `"internal"` - For internal/private endpoints
- `"external"` - For external/public endpoints

## Migration Guide

If you have existing config files with numeric values:

1. **Language**: Replace `language: 0` with `language: "go"`
2. **Placement**: Replace `placement: 1` with `placement: "external"`
3. **REST API Config**: Remove the `routes:` wrapper:
   ```yaml
   # Old:
   rest_api_config:
     routes:
       - route: "/path"
   
   # New:
   rest_api_config:
     - route: "/path"
   ```

## Files Updated

### In types package:
- ✅ `plugin.proto` - Updated schema
- ✅ `protobuff/plugin.pb.go` - Regenerated
- ✅ `inject_yaml_tags.go` - YAML tag injection script
- ✅ `Makefile` - Automated generation

### In engine package:
- ✅ `open-core/services/plugin/yaml_plugin_loader.go` - Updated loader logic

### In udbhabon-plugins:
- ✅ `hc-suchok-plugin/config.yml` - Updated to use string values

## Testing

The schema has been tested and verified to work correctly with YAML unmarshaling. All fields are properly populated and readable.

