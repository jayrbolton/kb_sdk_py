package cmd

import (
  "github.com/spf13/cobra"
  "log"
)

var rootCmd = &cobra.Command{
  Use: "kbase-sdk",
  Short: "Write, manage, and test KBase modules",
  Long: "Write, manage, and test KBase modules",
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("For help, run: kbase-sdk --help")
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
