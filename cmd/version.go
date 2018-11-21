package cmd
import (
  "fmt"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use: "version",
  Short: "Print version number",
  Long: "Print version number",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("KBase SDK CLI v0.0.1")
  },
}
