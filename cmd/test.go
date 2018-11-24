package cmd

import (
  "github.com/spf13/cobra"
  "fmt"
)

var testCmd = &cobra.Command{
  Use: "test",
  Short: "Test a KBase module",
  Long: "Test a KBase module",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("TODO validate the module")
    // - read in config files and validate
    // - build and tag if any changes to Dockerfile
    // - docker run test 
  },
}

func init() {
  rootCmd.AddCommand(testCmd)
}
