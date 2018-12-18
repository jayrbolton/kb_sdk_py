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
    // Write kbase_methods.yaml
    write_template(module_name + "/kbase_methods.yaml", kbase_methods_yaml, &empty_map)
    // Write the Dockerfile
    write_template(module_name + "/Dockerfile", dockerfile, &empty_map)
    // Write gitignore
    write_template(module_name + "/.gitignore", gitignore, &empty_map)
    // Write README.md
    write_template(module_name + "/README.md", readme, &module_info)
    // Write compile_report.json
    write_template(module_name + "/compile_report.json", compile_report, &module_info)
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


// Template for kbase.yaml
// Requires module name and username
var kbase_yaml = `module-name: {{.Name}}
module-description: Enter a description here
service-language: python
module-version: 0.0.1
owners: ["{{.Username}}"]
` // end kbase_yaml


// Template for kbase_methods.yaml
var kbase_methods_yaml = `echo:
  title: Echo
  description: Repeat a message back to you
  required_params: [message]
  params:
    message:
      title: Message
      decription: String to echo back
      type: string
      minLength: 1
` // end kbase_methods_yaml


// Template for the Dockerfile
var dockerfile = `FROM python:3.7-alpine

# Install pip dependencies
RUN apk --update add --virtual build-dependencies python-dev build-base && \
    pip install --upgrade --no-cache-dir --extra-index-url https://pypi.anaconda.org/kbase/simple \
      kbase_module \
      flake8 && \
    apk del build-dependencies

# Run the app
WORKDIR /kb/module
COPY . /kb/module
RUN chmod -R a+rw /kb/module
EXPOSE 5000
ENTRYPOINT ["sh", "/usr/local/bin/entrypoint.sh"]  # from the kbase_module package
` // end dockerfile


// Gitignore tempate with common defaults:
var gitignore = `
# KBase build files
compile_report.json

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

var readme = `# {{.Name}} (KBase module)

This is a [KBase](https://kbase.us) module. You need to install the SDK command-line interface to run this module.

<TODO -- enter an extended description here>

## Development

1. Install the [KBase SDK](https://github.com/jayrbolton/kbase_sdk_cli)
2. Run tests with: kbase-sdk test
`

var compile_report = `{
  "function_places": {},
  "functions": {},
  "impl_file_path": "",
  "module_name": "{{.Name}}",
  "sdk_git_commit": "0",
  "sdk_version": "0",
  "spec_files": [{"content": "", "is_main": 1}]
}
` // end compile_report
