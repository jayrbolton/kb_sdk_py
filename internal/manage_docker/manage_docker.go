package manage_docker

// Utilities for building and running containers

import (
  "fmt"
  "log"
  "github.com/jayrbolton/kbase_sdk_cli/internal/shell"
  "path/filepath"
  "os"
  "os/exec"
)

// Build the docker container
func Build (use_cache bool, module_name string) {
  build_args := []string{"build", ".", "-t", get_tag(module_name), "--build-arg", "DEVELOPMENT=1"}
  if !use_cache {
    build_args = append(build_args, "--no-cache")
  }
  log.Println("Building docker container...")
  shell.RunCommand("docker", build_args...)
}

// Run a command in the docker container
func RunCommand (command string, module_name string) {
  abs_path, _ := filepath.Abs("./")
  docker_tag := get_tag(module_name)
  // Build the container if the tag is not found
  container_exists := check_docker_tag(docker_tag)
  if !container_exists { Build(true, module_name) }
  mount_arg := fmt.Sprintf("%v:/kb/module", abs_path)
  docker_args := []string{"run", "-t", "-v", mount_arg}
  // Check for a .env file and pass it to docker if it exists
  env_path := ".env"
  _, err := os.Stat(env_path)
  if err == nil {
    docker_args = append(docker_args, "--env-file", env_path)
  }
  docker_args = append(docker_args, docker_tag, command)
  shell.RunCommand("docker", docker_args...)
}

// Get the docker tag name from the module name
func get_tag (module_name string) (string) {
 return fmt.Sprintf("kbase_modules/%v", module_name)
}

// Check whether a docker image tag exists already
func check_docker_tag (name string) bool {
  out, err := exec.Command("docker", "images", "-q", name).Output()
  if err != nil {
    log.Fatal(err)
  }
  return len(out) > 0
}
