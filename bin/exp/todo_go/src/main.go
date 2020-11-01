package main

import (
  "fmt"
  . "todo_go/src/fp"
  . "todo_go/src/util"
)


func main() {
//  result := Cons(1, -2, 3, -4, -5).Map(func(e Any) Any {
//    return e.(int) * 100
//  })//.FilterNot(func(e Any) bool {
//    //return e.(int) >= 0
//  //})
//
//  result.Foreach(func(e Any){
//    fmt.Println(-e.(int))
//  })

 // res := result.Exists(func(e Any) bool {
 //   return e.(int) == 300
 // })

 lines := FileLines("/home/temp/doc/worklog/current.md")
 lines.Foreach(func(line Any) {
   fmt.Println(line)
 })

}
