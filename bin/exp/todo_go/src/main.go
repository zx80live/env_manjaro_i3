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
  fmt.Println(TodoSelector)
 
  FileLines(filename).
    Filter(TodoSelector).
    ZipWithIndex().
    Foreach(func(e Any) {
      //t := e.(Tuple2)
      fmt.Println(e)
    })
}
