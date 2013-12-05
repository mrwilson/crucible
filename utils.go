package crucible

import (
  "os/user"
  "log"
)

const (
  crucibleDir    = ".crucible"
  ansibleCommand = "ansible-playbook"
)

func HandleErr(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func HomeDir() string {
  usr, err := user.Current(); HandleErr(err);

  return usr.HomeDir
}
