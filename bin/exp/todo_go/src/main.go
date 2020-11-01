package main

import (
  "fmt"
//  "regexp"
//  "os"
  . "todo_go/src/fp"
//  . "todo_go/src/util"
   "time"
)


func main() {
//  args := os.Args[1:]
//  filename := args[0]
//  regex := args[1]
// 
//  pattern := regexp.MustCompile(regex)
// 
//  TodoSelector := func(line Any) bool {
//    return pattern.Match([]byte(line.(string)))
//  }
//  fmt.Println(TodoSelector)
//
//  FileLines(filename).
//    Filter(TodoSelector).
//    Zip(indexes).
//    Foreach(func(e Any) {
//      //t := e.(Tuple2)
//      fmt.Println(e)
//    })

  logger := func(e Any) {
    fmt.Println(e)
  }

//  l1 := Cons(0,1,2,3,4,5).Map(func(e Any) Any {
//    time.Sleep(1 * time.Nanosecond)
//    return e
//  })
  l2 := Cons("zero", "one", "two", "three", "four", "five").Map(func(e Any) Any {
    time.Sleep(2 * time.Nanosecond)
    return e
  })

 // l1.Foreach(logger)
//  l2.Foreach(logger)
  fmt.Println("---")

 // l3 := l1.Zip(l2)
 // l3.Foreach(logger)
// res := l1.Fold(0, func(e Any, acc Any) Any {
//   newAcc := e.(int) + acc.(int)
//   return newAcc 
// })
res := l2.Fold("", func(e Any, acc Any) Any {
  return e.(string) + acc.(string)
})

 println(logger, "res=", res.(string))

  
}
