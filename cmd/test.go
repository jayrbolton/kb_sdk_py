package cmd

// Command for running python tests within the docker container

import (
  "github.com/spf13/cobra"
  "os/exec"
  "log"
  "github.com/jayrbolton/kbase_sdk_cli/internal/module_config"
  "github.com/jayrbolton/kbase_sdk_cli/internal/manage_docker"
)

func init() {
  rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
  Use: "test",
  Short: "Run tests on a KBase module",
  Run: func(cmd *cobra.Command, args []string) {
    command := exec.Command("docker", "--version")
    if err := command.Run(); err != nil {
      log.Fatal("Unable to run docker, make sure it is installed first: https://docs.docker.com/install/")
    }
    // Validate presence of module files
    module_config.CheckFiles()
    module, err := module_config.LoadModule()
    if err != nil { log.Fatal(err) }
    manage_docker.RunCommand("test", module.Name)
  },
}
