package crucible

import (
  "fmt"
  "os/exec"
  "path/filepath"
)

func RunPlaybook(options []string) {

  options[0] = filepath.Join(HomeDir(), crucibleDir, options[0], "site.yml")

  cmd := exec.Command(ansibleCommand, options...)
  cmd.Env = nil

  out, err := cmd.CombinedOutput(); HandleErr(err);

  fmt.Println(string(out))

}
