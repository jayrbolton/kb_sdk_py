package cmd
import (
  "log"
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/spf13/cobra"
  "github.com/jayrbolton/kbase_sdk_cli/internal/versioning"
  "github.com/jayrbolton/kbase_sdk_cli/internal/shell"
)

func init() {
  rootCmd.AddCommand(upgradeCmd)
}

var upgradeCmd = &cobra.Command{
  Use: "upgrade",
  Short: "Upgrade the KBase SDK CLI",
  Run: func(cmd *cobra.Command, args []string) {
    // Check the version on github
    new_version := versioning.Fetch()
    log.Printf("Current version: %v\n", versioning.CurrentVersion)
    log.Printf("Newest version: %v\n", new_version)
    install_url := fmt.Sprintf(
      "https://github.com/jayrbolton/kbase_sdk_cli/releases/download/%v/install.sh",
      versioning.CurrentVersion)
    // Check if the new version is greater than the current version
    if versioning.IsGreater(new_version, versioning.CurrentVersion) {
      log.Println("Upgrading to the newest version...")
      resp, err := http.Get(install_url)
      if err != nil {
        log.Fatalf("Error downloading installation script: %v\n", err)
      }
      defer resp.Body.Close()
      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        log.Fatalf("Error downloading installation script: %v\n", err)
      }
      shell.RunCommand("/bin/sh", "-c", string(body))
    } else {
      log.Println("Your CLI is up to date.")
    }
  },
}
