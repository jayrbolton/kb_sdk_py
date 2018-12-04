package cmd

import (
  "github.com/spf13/cobra"
  "os/exec"
  "os"
  "log"
  "fmt"
  "io/ioutil"
  "path/filepath"
  "github.com/jayrbolton/kbase_sdk_cli/internal/shell"
  "gopkg.in/yaml.v2"
)

// Flag -- whether to rebuild docker
var build bool
// Flag -- whether to rebuild docker without any caching
var build_no_cache bool

func init() {
  testCmd.PersistentFlags().BoolVar(&build, "build", false, "Rebuild the docker container before running tests.")
  testCmd.PersistentFlags().BoolVar(&build_no_cache, "build-no-cache", false, "Rebuild the docker container with no cache.")
  rootCmd.AddCommand(testCmd)
}

type Module struct {
  Name string `yaml:"module-name"`
}

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
    check_for_file("./kbase.yaml")
    check_for_file("./src/main.py")
    check_for_file("./Dockerfile")
    // Read in the module name from kbase.yaml
    var module Module
    kbase_module_bytes, err := ioutil.ReadFile("kbase.yaml")
    if err != nil {
      log.Fatalf("Unable to open kbase.yaml: %v\n", err)
    }
    err = yaml.Unmarshal(kbase_module_bytes, &module)
    if err != nil {
      log.Fatalf("Unable to decode kbase.yaml: %v\n", err)
    }
    log.Printf("Module name: %v\n", module.Name)
    docker_tag := fmt.Sprintf("kbase_modules/%v", module.Name)
    log.Printf("Docker tag: %v\n", docker_tag)
    log.Println("To re-build your docker container, run: kbase-sdk test --build")
    log.Printf("To open a shell in your container, run: docker run -it %v shell\n", docker_tag)
    // Build the image if the tag is not found
    image_exists := check_docker_tag(docker_tag)
    if !image_exists {
      build = true
    }
    // build is global to this file, set either as a flag or above
    if build {
      log.Println("Building docker container...")
      shell.RunCommand("docker", "build", ".", "-t", docker_tag)
    } else if build_no_cache {
      log.Println("Building docker container without any caching...")
      shell.RunCommand("docker", "build", ".", "-t", docker_tag, "--no-cache")
    }
    log.Println("Running tests...")
    mount_arg := fmt.Sprintf("%v:/kb/module", abs_path)
    // Uses the entrypoint.sh script from the kbase_module package
    shell.RunCommand("docker", "run", "-v", mount_arg, docker_tag, "test")
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

func check_for_file (path string) {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    log.Fatalf("%v does not exist. Are you in a KBase module directory?", path)
  }
}
