package main

import (
  "github.com/mrwilson/crucible"
  "os"
)

func main() {

  if os.Args[1] == "get" {
    crucible.DownloadPlayBook(os.Args[2])
  } else {
    crucible.RunPlaybook(os.Args[1:])
  }

}

