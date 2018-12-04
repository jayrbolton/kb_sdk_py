package cmd

import (
  "github.com/spf13/cobra"
  "os"
  "regexp"
  "text/template"
  "log"
)

// Add the init command to the set of commands in root
func init() {
  rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
  Use: "init [module name]",
  Short: "Initialize a new KBase SDK module",
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    var module_name = args[0]
    // Check for the necessary env vars
    kbase_username := os.Getenv("KBASE_USERNAME")
    if kbase_username == "" {
      log.Fatal("Please set the KBASE_USERNAME environment variable first.\n")
    }
    // Check format of the project name
    match, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z_\-0-9]+$`, module_name)
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
    // Write kbase.yaml
    module_info := map[string]string{"Name": module_name, "Username": kbase_username}
    write_template(module_name + "/kbase.yaml", kbase_yaml, &module_info)
    // Write kbase_methods.json
    write_template(module_name + "/kbase_methods.json", kbase_methods_json, &empty_map)
    // Write the Dockerfile
    write_template(module_name + "/Dockerfile", dockerfile, &empty_map)
    // Write requirements.txt
    write_template(module_name + "/requirements.txt", requirements_txt, &empty_map)
    // Write gitignore
    write_template(module_name + "/.gitignore", gitignore, &empty_map)
    // We're done.
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


def echo(params):
    """Echo back the given message."""
    return params['message']
` // end main_py


// Template for src/test/test_main.py
var test_main_py = `"""
This file contains tests for src/main.py
"""
import kbase_module
import unittest


class TestMain(unittest.TestCase):

    def test_echo(self):
        """Test the echo function."""
        message = "Hello world!"
        result = kbase_module.run_method('echo', {'message': message})
        self.assertEqual(result, message)

    def test_echo_invalid_params(self):
        """Test the case where we don't pass the 'message' param."""
        with self.assertRaises(RuntimeError) as context:
            kbase_module.run_method('echo', {})
        msg = "'message' is a required property"
        self.assertTrue(msg in str(context.exception))
` // end test_main_py


// Template for src/kbase-module.json
// Requires module name and username
var kbase_yaml = `module-name: {{.Name}}
module-description: Enter a description here
service-language: python
module-version: 0.0.1
owners: ["{{.Username}}"]
` // end kbase_yaml


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


// Template for the Dockerfile
var dockerfile = `FROM python:3.7-alpine

# Install pip dependencies
WORKDIR /kb/module
COPY requirements.txt /kb/module/requirements.txt
RUN apk --update add --virtual build-dependencies python-dev build-base && \
    pip install --upgrade --no-cache-dir pip -r requirements.txt && \
    apk del build-dependencies

# Run the app
COPY . /kb/module
RUN chmod -R a+rw /kb/module
EXPOSE 5000
ENTRYPOINT ["sh", "/usr/local/bin/entrypoint.sh"]  # from the kbase_module package
` // end dockerfile


// Template for the pip dependencies in requirements.txt
var requirements_txt = `--extra-index-url https://pypi.anaconda.org/kbase/simple
# This is needed for basic module functionality 
kbase_module
# For checking Pep8 syntax standards
flake8>3` // end requirements_txt


// Gitignore tempate with common defaults:
var gitignore = `
# Byte-compiled / optimized / DLL files
__pycache__/
*.py[cod]
*$py.class

# C extensions
*.so

# Distribution / packaging
eggs/
.eggs/
*.egg-info/
*.egg
MANIFEST

# Installer logs
pip-log.txt
pip-delete-this-directory.txt

# Environments
.env
.venv
env/
venv/
ENV/
env.bak/
venv.bak/

# mypy
.mypy_cache/
.dmypy.json
dmypy.json
` // end gitignore
