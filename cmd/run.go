package cmd

// Command for running a method in the module with some json input

import (
  "github.com/spf13/cobra"
  "os/exec"
  "log"
  "github.com/jayrbolton/kbase_sdk_cli/internal/module_config"
  "github.com/jayrbolton/kbase_sdk_cli/internal/manage_docker"
)

// Add the init command to the set of commands in root
func init() {
  rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
  Use: "run",
  Short: "Run a module method with some input json",
  Run: func(cmd *cobra.Command, args []string) {
    command := exec.Command("docker", "--version")
    if err := command.Run(); err != nil {
      log.Fatal("Unable to run docker, make sure it is installed first: https://docs.docker.com/install/")
    }
    // Validate presence of module files
    module_config.CheckFiles()
    module, err := module_config.LoadModule()
    if err != nil { log.Fatal(err) }
    // TODO json from a file `kbase-sdk run [method name] -f input.json`
    // TODO get the method name from the arg: `kbase-sdk run [method name] '{}'`
    // TODO place the input in /kb/module/work/input.json
    manage_docker.RunCommand("async", module.Name)
    // TODO read the output from /kb/module/work/input.json
  },
}
