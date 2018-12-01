package shell

import (
  "log"
  "os/exec"
  "bufio"
)

// Run a command, logging all output
func RunCommand (name string, arg ...string) {
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
