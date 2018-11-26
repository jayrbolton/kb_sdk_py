package cmd
import (
  "fmt"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
  Use: "status",
  Short: "View module status in the KBase catalog.",
  Long: "View module status in the KBase catalog.",
  Run: func(cmd *cobra.Command, args []string) {
    // TODO get this from some config or env
    fmt.Println("TODO show status of this module in the catalog.")
  },
}

