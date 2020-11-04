package main

import "fmt"
import "os"
import "regexp"

func main() {
  args := os.Args[1:]
  if len(args) < 2 {
    return;
  }

  text := args[0]
  regex := args[1]

  r := regexp.MustCompile(regex)
  for _, e := range r.FindStringSubmatch(text) {
    fmt.Printf("%v\n", e)
  }
}
