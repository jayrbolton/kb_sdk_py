package cmd

import (
  "github.com/spf13/cobra"
  "fmt"
  "os"
)

var rootCmd = &cobra.Command{
  Use: "kbase-sdk",
  Short: "Write, manage, and test KBase modules",
  Long: "Write, manage, and test KBase modules",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("For help, run: kbase-sdk --help")
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
