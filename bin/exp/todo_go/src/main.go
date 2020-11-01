package main

import (
  "fmt"
  "regexp"
  "os"
  . "todo_go/src/fp"
  . "todo_go/src/util"
)


func main() {
  args := os.Args[1:]
  filename := args[0]
  regex := args[1]
 
  pattern := regexp.MustCompile(regex)
 
  TodoSelector := func(line Any) bool {
    return pattern.Match([]byte(line.(string)))
  }
 
  FileLines(filename).
    Filter(TodoSelector).
    Foreach(func(e Any) {
      fmt.Println("> " + e.(string))
    })
}
