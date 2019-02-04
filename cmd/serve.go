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
  rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
  Use: "serve",
  Short: "Run a module as a JSON RPC server",
  Run: func(cmd *cobra.Command, args []string) {
    command := exec.Command("docker", "--version")
    if err := command.Run(); err != nil {
      log.Fatal("Unable to run docker, make sure it is installed first: https://docs.docker.com/install/")
    }
    // Validate presence of module files
    module_config.CheckFiles()
    module, err := module_config.LoadModule()
    if err != nil { log.Fatal(err) }
    manage_docker.RunCommand("serve", module.Name)
  },
}

