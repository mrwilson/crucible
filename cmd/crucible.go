package main

import (
  "github.com/mrwilson/crucible"
  "os"
)

func main() {
  crucible.RunPlaybook(os.Args[1:])
}

