package main

import "fmt"
import . "todo_go/src/fp"

func Pos(e Any) bool {
  return e.(int) >= 0
}

func Neg(e Any) bool {
  return e.(int) < 0
}

func Even(e Any) bool {
  return e.(int) % 2 == 0
}

func Odd(e Any) bool {
  return e.(int) % 2 != 0
}

func Log(e Any) {
  fmt.Println(e)
}

func Mul100(e Any) Any {
  return e.(int) * 100
}

func Sum(a Any, b Any) Any {
  return a.(int) + b.(int)
}

func StrConcat(a Any, b Any) Any {
  return a.(string) + b.(string)
}

func main() {
  fmt.Println("Hello")

//  l := Nil.Cons(10).Cons(9).Cons(8).Cons(7).Cons(-6).Cons(5).Cons(-4).Cons(3).Cons(2).Cons(1).Cons(0)
//
//  l.
//    Filter(Even).
//    Filter(Pos).
//    Map(Mul100).
//    Foreach(Log)
//
//  fmt.Println("--")
//  l.Filter(Neg).Foreach(Log)
//
//  fmt.Println("--")
//  l.Filter(Odd).Foreach(Log)

  fmt.Println("--")
  l1 := Nil.Cons(6).Cons(5).Cons(4).Cons(3).Cons(2).Cons(1).Cons(0)
  l2 := Nil.Cons("f").Cons("e").Cons("d").Cons("c")
  l3 := Nil.Cons("fff").Cons("eee").Cons("ddd").Cons("ccc").Cons("bbb").Cons("aaa")
  l1.Foreach(Log)
  l2.Foreach(Log)
  l1.Zip(l2).Foreach(Log)
  fmt.Println("--")
  l1.Zip(l2).Zip(l3).Foreach(Log)
  fmt.Println("--")
  l2.ZipWithIndex().Foreach(Log)
  fmt.Println("--")
  fmt.Println("even", l1.Count(Even), "odd", l1.Count(Odd), "neg", l1.Count(Neg), "pos", l1.Count(Pos))
  fmt.Println("size", l1.Size(), "Nil.size=", Nil.Size(), "tail.size=",  l1.Tail().Size())
  fmt.Println("5 exist", l1.Exist(func(e Any) bool {return e.(int) == 5}))
  fmt.Println("100 exist", l1.Exist(func(e Any) bool {return e.(int) == 100}))
  fmt.Println("even esist", l1.Exist(Even))
  fmt.Println("--- reduce ---")
  fmt.Println(l1.Reduce(Sum))
  fmt.Println(l2.Reduce(StrConcat))
  fmt.Println("--- fold left ---")
  res1 := l1.FoldLeft(1000, func(acc Any, e Any) Any {
    return acc.(int) + e.(int)
  })
  fmt.Println(res1)
}
