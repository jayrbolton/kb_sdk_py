package shell

import (
  "log"
  "os/exec"
  "os"
)

// Run a command, logging all output
func RunCommand (name string, arg ...string) {
  cmd := exec.Command(name, arg...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Start()
  err := cmd.Wait()
  if err != nil { fatal(name, err) }
}

func fatal (name string, err error) {
  log.Fatalf("Unable to run %v (see logs above)\n%v", name, err)
}
