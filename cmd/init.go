package cmd

import (
  "github.com/spf13/cobra"
  "fmt"
)

var initCmd = &cobra.Command{
  Use: "init [module name]",
  Short: "Initialize a new KBase module",
  Long: "Initialize a new KBase module",
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("TODO initialize the things!")
    // - check for and create project directory
    // - create src/project and src/test directories
    // - create src/project/main.py
    // - create src/test/test_main.py
    // - create kbase-module.json and kbase-methods.json with hello world example
  },
}

func init() {
  rootCmd.AddCommand(initCmd)
}
