package crucible

import (
  "fmt"
  "os/exec"
  "strings"
)

const (
  ansibleCommand string = "ansible-playbook"
  crucibleDir    string = "/home/mrwilson/.crucible"
)

func RunPlaybook(options []string) {

  options[0] = strings.Join([]string{crucibleDir,options[0],"site.yml"}, "/")

  cmd := exec.Command(ansibleCommand, options...)
  cmd.Env = nil
  out, err := cmd.CombinedOutput()

  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(out))

}
