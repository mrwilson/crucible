package crucible

import (
  "fmt"
  "os/exec"
  "os/user"
  "strings"
)

const (
  ansibleCommand string = "ansible-playbook"
  crucibleDir    string = ".crucible"
)

func RunPlaybook(options []string) {

  usr, err := user.Current()

  if err != nil {
    fmt.Println(err)
  }

  playBookPath := []string{ usr.HomeDir, crucibleDir, options[0], "site.yml" }

  options[0] = strings.Join(playBookPath, "/")

  cmd := exec.Command(ansibleCommand, options...)
  cmd.Env = nil
  out, err := cmd.CombinedOutput()

  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(out))

}
