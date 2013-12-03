package main

import (
  "fmt"
  "github.com/mrwilson/crucible"
  "os"
)

func main() {
  fmt.Println(os.Args[1:])
  crucible.RunPlaybook(os.Args[1:])
}

