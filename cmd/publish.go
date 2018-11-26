package cmd
import (
  "fmt"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(publishCmd)
}

var publishCmd = &cobra.Command{
  Use: "status",
  Short: "Publish this module to the KBase catalog.",
  Long: "Publish this module to the KBase catalog.",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("TODO publish an update to the catalog.")
  },
}


