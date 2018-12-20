package shell

import (
  "log"
  "os/exec"
  "bufio"
)

// Run a command, logging all output
func RunCommand (name string, arg ...string) {
  cmd := exec.Command(name, arg...)
  stdout, _ := cmd.StdoutPipe()
  scanner_out := bufio.NewScanner(stdout)
  stderr, _ := cmd.StderrPipe()
  scanner_err := bufio.NewScanner(stderr)
  cmd.Start()
  go func () {
    for scanner_out.Scan() {
      log.Print(scanner_out.Text())
    }
  }()
  go func () {
    for scanner_err.Scan() {
      log.Print(scanner_err.Text())
    }
  }()
  err := cmd.Wait()
  if err != nil { fatal(name, err) }
}

func fatal (name string, err error) {
  log.Fatalf("")
  // log.Fatalf("Unable to run %v (see logs above)\n%v", name, err)
}
