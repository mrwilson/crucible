package crucible

import (
  "fmt"
  "os/exec"
)

const ansibleCommand string = "ansible-playbook"

func RunPlaybook(options []string) {

  cmd := exec.Command(ansibleCommand, options...)
  cmd.Env = nil
  out, err := cmd.CombinedOutput()

  if err != nil {
    fmt.Println("It's all gone wrong!")
  }

  fmt.Println(string(out))
}
