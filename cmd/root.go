package cmd

import (
  "github.com/spf13/cobra"
  "log"
  "github.com/jayrbolton/kbase_sdk_cli/internal/versioning"
)

var rootCmd = &cobra.Command{
  Use: "kbase-sdk",
  Short: "Write, manage, and test KBase modules. Show CLI version with --version.",
  Version: versioning.CurrentVersion,
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("For help, run: kbase-sdk --help")
    new_version := versioning.Fetch()
    if versioning.IsGreater(new_version, versioning.CurrentVersion) {
      log.Printf("There is new version of the CLI: %v\n", new_version)
      log.Println("Run \"kbase-sdk upgrade\" to upgrade to the newest version.")
    }
  },
}

func init () {
  // Don't log timestamps
  log.SetFlags(0)
}

func Execute () {
  if err := rootCmd.Execute(); err != nil {
    log.Fatal(err)
  }
}
