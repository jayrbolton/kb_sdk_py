package cmd

import (
  "github.com/spf13/cobra"
  "os/exec"
  "os"
  "log"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "bufio"
  "path/filepath"
)

func init() {
  testCmd.PersistentFlags().BoolVar(&build, "build", false, "Rebuild the docker container before running tests.")
  testCmd.PersistentFlags().BoolVar(&build_no_cache, "build-no-cache", false, "Rebuild the docker container with no cache.")
  rootCmd.AddCommand(testCmd)
}

type Module struct {
  Name string `json:"name"`
}

// Flag -- whether to rebuild docker
var build bool
// Flag -- whether to rebuild docker without any caching
var build_no_cache bool

var testCmd = &cobra.Command{
  Use: "test",
  Short: "Run tests on a KBase module",
  Run: func(cmd *cobra.Command, args []string) {
    command := exec.Command("docker", "--version")
    if err := command.Run(); err != nil {
      log.Fatal("Unable to run docker, make sure it is installed first: https://docs.docker.com/install/")
    }
    abs_path, err := filepath.Abs("./")
    if err != nil {
      log.Fatal(err)
    }
    // Check for presence of basic config files
    check_for_file("./kbase_methods.json")
    check_for_file("./kbase_module.json")
    check_for_file("./src/main.py")
    check_for_file("./Dockerfile")
    // Read in the module name from kbase_module.json
    var module Module
    kbase_module_bytes, err := ioutil.ReadFile("kbase_module.json")
    if err != nil {
      log.Fatalf("Unable to open kbase_module.json: %v", err)
    }
    json.Unmarshal(kbase_module_bytes, &module)
    log.Printf("Module name: %v", module.Name)
    docker_tag := fmt.Sprintf("kbase_modules/%v", module.Name)
    // Build the image if the tag is not found
    image_exists := check_docker_tag(docker_tag)
    if !image_exists {
      build = true
    }
    if build {
      log.Println("Building docker container...")
      run_command("docker", "build", ".", "-t", docker_tag)
    } else if build_no_cache {
      log.Println("Building docker container without any caching...")
      run_command("docker", "build", ".", "-t", docker_tag, "--no-cache")
    }
    log.Println("Running tests...")
    mount_arg := fmt.Sprintf("%v:/kb/module", abs_path)
    run_command("docker", "run", "-v", mount_arg,
      docker_tag, "python", "-m", "unittest", "discover", "/kb/module/src/test")
  },
}

// Check whether a docker image tag exists already
func check_docker_tag (name string) bool {
  out, err := exec.Command("docker", "images", "-q", name).Output()
  if err != nil {
    log.Fatal(err)
  }
  return len(out) > 0
}

// Run a command, logging all output
func run_command (name string, arg ...string) {
  cmd := exec.Command(name, arg...)
  stdout, _ := cmd.StderrPipe()
  stderr, _ := cmd.StdoutPipe()
  cmd.Start()
  scanner_err := bufio.NewScanner(stderr)
  scanner_out := bufio.NewScanner(stdout)
  for scanner_err.Scan() {
    log.Print(scanner_err.Text())
  }
  for scanner_out.Scan() {
    log.Print(scanner_out.Text())
  }
  err := cmd.Wait()
  if err != nil {
    log.Fatal("Unable to build container (see above)")
  }
}

func check_for_file (path string) {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    log.Fatalf("%v does not exist. Are you in a KBase module directory?", path)
  }
}
