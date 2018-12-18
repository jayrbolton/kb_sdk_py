package cmd

// Command for building files necessary for registration on the catalog
// In particular, everything under /ui/etc

import (
  // "log"
  "github.com/spf13/cobra"
  // "github.com/jayrbolton/kbase_sdk_cli/internal/module_config"
)

// Add the init command to the set of commands in root
func init() {
  rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
  Use: "build",
  Short: "Generate build files for a KBase SDK module",
  Run: func(cmd *cobra.Command, args []string) {
    // module, err := module_config.LoadModule()
    // if err != nil {
    //   log.Fatalf("Unable to load kbase.yaml: %s\n", err)
    // }
    // methods := module_config.LoadMethods()
    // log.Printf("Module: %s\n", module)
    // log.Printf("Methods:\n%s\n", methods)
    // module_config.WriteCompileReport(&module, &methods)
    // TODO WriteUIConfig(&module, &methods)
  },
}
