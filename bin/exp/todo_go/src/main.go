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

//  logger := func(e Any) {
//    fmt.Println(e)
//  }

  l1 := Cons(0,1,2,3,4,5,6,7,8,9,10,11,12, 13,14).Map(func(e Any) Any {
    time.Sleep(1 * time.Nanosecond)
    return e
  })
//  l2 := Cons("t0", "t1", "t2", "t3", "t4", "t5", "t6", "t7", "t8", "t9", "t10", "t11", "t12", "t13").Map(func(e Any) Any {
//    time.Sleep(5000 * time.Nanosecond)
//    return e
//  })

 // l3 := Cons("eet0", "eet1", "eet2", "eet3", "eet4", "eet5", "eet6", "eet7", "eet8", "eet9", "eet10", "eet11", "eet12", "eet13").Map(func(e Any) Any {
 //   time.Sleep(5 * time.Nanosecond)
 //   return e
 // })

 // l1.Foreach(logger)
//  l2.Foreach(logger)
  fmt.Println("---")

 //l4 := l1.ZipWith(l2)
 Zip22(l1)
// l4.Foreach(logger)
// res := l1.Fold(0, func(e Any, acc Any) Any {
//   newAcc := e.(int) + acc.(int)
//   return newAcc 
// })
//  res := l2.Fold("", func(e Any, acc Any) Any {
//    return e.(string) + acc.(string)
//  })
//
//   println(logger, "res=", res.(string))


}
