package cmd

// Command for building files necessary for registration on the catalog
// In particular, everything under /ui/etc

import (
  "github.com/spf13/cobra"
  "os/exec"
  "log"
  "github.com/jayrbolton/kbase_sdk_cli/internal/module_config"
  "github.com/jayrbolton/kbase_sdk_cli/internal/manage_docker"
)

var no_cache bool

// Add the init command to the set of commands in root
func init() {
  buildCmd.PersistentFlags().BoolVar(&no_cache, "no-cache", false,
    "Build the docker container with no cache.")
  rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
  Use: "build",
  Short: "Build the docker container and auto-generate files for registration in the catalog.",
  Run: func(cmd *cobra.Command, args []string) {
    command := exec.Command("docker", "--version")
    if err := command.Run(); err != nil {
      log.Fatal("Unable to run docker, make sure it is installed first: https://docs.docker.com/install/")
    }
    // Validate presence of module files
    module_config.CheckFiles()
    module, err := module_config.LoadModule()
    if err != nil { log.Fatal(err) }
    manage_docker.Build(!no_cache, module.Name)
  },
}
