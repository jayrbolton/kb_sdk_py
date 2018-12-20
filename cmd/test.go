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
  "github.com/jayrbolton/kbase_sdk_cli/internal/module_config"
  "gopkg.in/yaml.v2"
)

// Flag -- whether to rebuild docker
var build bool
// Flag -- whether to rebuild docker without any caching
var build_no_cache bool

func init() {
  testCmd.PersistentFlags().BoolVar(&build, "build", false,
    "Rebuild the docker container before running tests.")
  testCmd.PersistentFlags().BoolVar(&build_no_cache, "build-no-cache", false,
    "Rebuild the docker container with no cache.")
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
    // Validate presence of module files
    module_config.CheckFiles()
    if err != nil {
      log.Fatal(err)
    }
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
    if build || build_no_cache {
      build_args := []string{"build", ".", "-t", docker_tag, "--build-arg", "DEVELOPMENT=1"}
      if build_no_cache {
        build_args = append(build_args, "--no-cache")
      }
      log.Println("Building docker container...")
      shell.RunCommand("docker", build_args...)
    }
    mount_arg := fmt.Sprintf("%v:/kb/module", abs_path)
    docker_args := []string{"run", "-v", mount_arg}
    // Check for a .env file and pass it to docker if it exists
    env_path := ".env"
    _, err = os.Stat(env_path)
    if err == nil {
      docker_args = append(docker_args, "--env-file", env_path)
    }
    docker_args = append(docker_args, docker_tag, "test")
    shell.RunCommand("docker", docker_args...)
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
