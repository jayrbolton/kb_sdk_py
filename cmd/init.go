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
  },
}

func init() {
  rootCmd.AddCommand(initCmd)
}
