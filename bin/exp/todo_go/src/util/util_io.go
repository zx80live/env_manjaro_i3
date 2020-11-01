package util

import . "todo_go/src/fp"
import "bufio"
import "os"
import "log"

func FileLines(filename string) Monad {
  c := make(Monad)

  go func(){
    defer close(c)

    file, err := os.Open(filename)
    defer file.Close()

    if err != nil {
      log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      c <- scanner.Text()
    }

    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }
  }()

  return c
}
