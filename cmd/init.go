package cmd

import (
  "github.com/spf13/cobra"
  "os"
  "regexp"
  "text/template"
  "log"
)

var initCmd = &cobra.Command{
  Use: "init [module name]",
  Short: "Initialize a new KBase module",
  Long: "Initialize a new KBase module",
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    var module_name = args[0]
    // Check for the necessary env vars
    kbase_username := os.Getenv("KBASE_USERNAME")
    if kbase_username == "" {
      log.Fatal("Please set the KBASE_USERNAME environment variable first.\n")
    }
    // Check format of the project name
    match, _ := regexp.MatchString(`^[a-zA-Z_\-0-9]+$`, module_name)
    if !match {
      log.Fatalf("Invalid module name \"%v\". It should be alphanumeric.\n", module_name)
    }
    log.Printf("Initializing new SDK module \"%v\"...\n", module_name)
    // Check for pre-existing project directory
    _, statErr := os.Stat(module_name)
    if !os.IsNotExist(statErr) {
      log.Fatalf("Directory already exists: %v/\n", module_name)
    } else {
      log.Println("Creating directories..")
      // Create project/, project/src/, and project/src/test/
      os.MkdirAll(module_name + "/src/test", os.ModePerm)
    }
    // We use an empty map for templates that need no data (yet)
    empty_map := make(map[string]string)
    // Write main.py
    write_template(module_name + "/src/main.py", main_py, &empty_map)
    // Write test_main.py
    write_template(module_name + "/src/test/test_main.py", test_main_py, &empty_map)
    // Write kbase_module.json
    module_info := map[string]string{"Name": module_name, "Username": kbase_username}
    write_template(module_name + "/kbase_module.json", kbase_module_json, &module_info)
    // Write kbase_methods.json
    write_template(module_name + "/kbase_methods.json", kbase_methods_json, &empty_map)
    log.Printf("Your new module lives in ./%v\n", module_name)
    log.Printf("Get started with: cd %v && kbase-sdk test\n", module_name)
  },
}

// Write a template out to a file with logging and error handling
func write_template(path string, templ_content string, config *map[string]string) {
  log.Printf("Writing %v..\n", path)
  templ, err := template.New(path).Parse(templ_content)
  if err != nil {
    log.Fatalf("Unable to create %v: %v", path, err)
  }
  file, err := os.Create(path)
  if err != nil {
    log.Fatalf("Unable to create file: %v\n", err.Error())
  }
  defer file.Close()
  templ.Execute(file, config)
}


// Template for src/main.py
var main_py = `"""
This file contains all the methods for this KBase SDK module.
"""
import kbase_module


@kbase_module.method('echo')
def echo(params):
    """Echo back the given message."""
    return params['message']
` // end main_py


// Template for src/test/test_main.py
var test_main_py = `"""
This file contains tests for src/main.py
"""
import unittest
from main import echo


class TestMain(unittest.TestCase):

    def test_echo(self):
        message = "Hello world!"
        result = echo(message)
        self.assertEqual(result, message)
` // end test_main_py


// Template for src/kbase-module.json
var kbase_module_json = `{
  "name": "{{.Name}}",
  "version": "0.0.1",
  "author": "{{.Username}}",
  "sdk_version": "2.0.0"
}
` // end kbase_module_json


// Template for src/kbase-methods.json
var kbase_methods_json = `{
  "echo": {
    "label": "Echo a message back to you.",
    "required_params": ["message"],
    "params": {
      "message": {"type": "string"}
    }
  }
}
` // end kbase_methods_json

// Add this command to the set of commands in root
func init() {
  rootCmd.AddCommand(initCmd)
}
