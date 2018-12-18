package module_config

import (
  "log"
  "os"
  "encoding/json"
  "io/ioutil"
  "github.com/xeipuuv/gojsonschema"
  "github.com/ghodss/yaml"
)

// Check that the current working directory is a valid KBase SDK Module
func CheckFiles () {
  check_for_file("./kbase_methods.yaml")
  check_for_file("./kbase.yaml")
  check_for_file("./src/main.py")
  check_for_file("./Dockerfile")
}

// Struct type for the kbase.yaml file
type Module struct {
  Name string `json:"module-name"`
  Language string `json:"service-language"`
  Owners []string `json:"owners"`
  Description string `json:"module-description"`
  Version string `json:"module-version"`
}

// Load and validate kbase.yaml
// We convert the YAML to JSON, validate it, then turn it into a struct
func LoadModule () (*Module, error) {
  json_config := validate_file("./kbase.yaml", kbase_yaml_schema)
  var module Module
  err := json.Unmarshal(json_config, &module)
  if err != nil { return nil, err }
  return &module, nil
}

// Struct for a kbase_methods.json entry
type Method struct {
  Label string `json:"label"`
  RequiredParams []string `json:"required_params"`
  // Possibly extremely nested json schema
  // {"paramName": {"type": "string", "properties": {etc...}}}
  Params map[string]interface{} // map[string]interface{}
}

// Validate and load kbase_methods.yaml
func LoadMethods () interface{} {
  json_config := validate_file("./kbase_methods.yaml", kbase_methods_yaml_schema)
  var methods map[string]Method
  err := json.Unmarshal(json_config, &methods)
  if err != nil {
    log.Fatalf("Error parsing kbase_methods.yaml:\n%v\n", err)
  }
  // Remove and recreate the ui/ directory
  os.RemoveAll("./ui/*")
  // We will always create some config files in ui/local_functions
  os.MkdirAll("./ui", 0700)
  // Iterate over each method
  for method_name, val1 := range methods {
    label := val1.Label
    required_params := val1.RequiredParams
    params := val1.Params
    // If the method has `narrative_app: true`, then we create ui/narrative/methods/<method_name>/etc
    // Otherwise, we create ui/local_functions/<method_name>.json
    log.Printf("Method: %v\n", method_name)
    log.Printf("  Label: %v\n", label)
    log.Printf("  Required: %v\n", required_params)
    log.Printf("  Params: %v\n", params)
    // For each method we create a ui/local_functions/<method_name>.json
    // for param_name, js_type := range params {
    //   type_name := js_type["type"]
    //   log.Printf("Param %v\n", param_name)
    //   log.Printf("  Type %v\n", type_name)
    // }
  }
  return methods
}

// Use JSON schema to validate a file
func validate_file (path string, schema string) []byte {
  file_bytes, err := ioutil.ReadFile(path)
  if err != nil {
    log.Fatalf("Unable to read %v:\n%v\n", path, err)
  }
  json_config, err := yaml.YAMLToJSON(file_bytes)
  if err != nil {
    log.Fatalf("Unable to parse %v:\n%v\n", path, err)
  }
  schema_loader := gojsonschema.NewStringLoader(schema)
  doc_loader := gojsonschema.NewStringLoader(string(json_config))
  result, err := gojsonschema.Validate(schema_loader, doc_loader)
  if err != nil {
    log.Printf("Unable to validate kbase.yaml file:\n%v\n", err)
  }
  if result.Valid() {
    log.Printf("%v looks valid\n", path)
  } else {
    log.Println("kbase.yaml validation errors:")
    for _, desc := range result.Errors() {
      log.Printf(" - %s\n", desc)
    }
    log.Fatalf("errors in the kbase.yaml file (above)\n")
  }
  return json_config
}

/*
// Convert data from kbase_methods.yaml into ui/.../spec.json and ui/.../display.yaml
// TODO
func CompileUIConfig (*map[string]Method) error {
  // TODO make a struct for spec.json and display.yaml
  // Iterate over methods and mutate spec_json and display_yaml
  // Convert the structs into json and write them out to their paths
  // We can leave the display_yaml as json since it is valid yaml
}
*/

// Check for the existence of a file relative to a path
// Logs and exits if it doesn't exist
func check_for_file (path string) {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    log.Fatalf("%v does not exist. Are you in a KBase module directory?", path)
  }
}

// JSON schema for kbase.yaml
var kbase_yaml_schema = `
{
  "type": "object",
  "required": ["module-name", "service-language", "owners", "module-version"],
  "properties": {
    "module-name": {
      "type": "string",
      "examples": ["megahit"],
      "pattern": "^[a-zA-Z_\\-][a-zA-Z_\\-0-9]+$",
      "minLength": 3
    },
    "owners": {
      "description": "KBase usernames",
      "type": "array",
      "examples": [["username1", "username2"]],
      "minItems": 1,
      "items": {"type": "string"}
    },
    "module-description": {"type": "string"},
    "module-version": {
      "description": "Semantic version",
      "type": "string",
      "examples": ["0.0.1", "v1.1.1"],
      "pattern": "v?\\d+\\.\\d+\\.\\d+"
    }
  }
}
`

// JSON schema for kbase_methods.yaml
var kbase_methods_yaml_schema = `
{
  "type": "object",
  "patternProperties": {
    "^.+$": {
      "type": "object",
      "properties": {
        "label": {"type": "string"},
        "description": {"type": "string"},
        "required_params": {
          "type": "array",
          "default": [],
          "items": {"type": "string"}
        },
        "params": {
          "default": {},
          "type": "object"
        }
      }
    }
  }
}
`
