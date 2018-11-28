package versioning

import (
  "log"
  "net/http"
  "strings"
  "io/ioutil"
)

var CurrentVersion = "v0.0.1"
var VersionURL = "https://raw.githubusercontent.com/jayrbolton/kbase_sdk_cli/master/VERSION"

// Fetch the latest version of the CLI from github
func Fetch () string {
  resp, err := http.Get(VersionURL)
  if err != nil {
    log.Fatalf("Unable to fetch latest version number: %v\n", err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  new_version := strings.TrimSuffix(string(body), "\n")
  return new_version
}
