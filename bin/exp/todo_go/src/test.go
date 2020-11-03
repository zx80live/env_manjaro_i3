package main

import "fmt"

type Any interface{}
type Predicate func(e Any) bool
type Functor func(e Any) Any

type Processor func(e Any) Any

type Traversable []Any
type Monad struct {
  content Traversable
  processor Processor
}

type Tuple2 struct {
  _1 Any
  _2 Any
}

var EmptyPredicate Predicate = func(e Any) bool { return true }
var EmptyFunctor Functor = func(e Any) Any { return e }
var EmptyProcessor Processor = func(e Any) Any {return e}

type List struct {
  Head Any
  Tail *List
}

var Nil List = List { nil, nil }

func (l List) Add(e Any) List {
  xs := List {
    Head: e,
    Tail: &l,
  }

  return xs
}

func (l List) Println() {
  if l.Head != nil {
    fmt.Println(l.Head)
  }

  if l.Tail != nil {
    l.Tail.Println()
  }
}

func (p1 Predicate) Filter(p2 Predicate) Predicate {
  res := func(e Any) bool {
    return p1(e) && p2(e)
  }
  return res
}

func (f1 Functor) Map(f2 Functor) Functor {
  res := func(e Any) Any {
    return f2(f1(e))
  }
  return res
}

func FSeq(els ... Any) Monad {
  return Monad {
    content: els,
    processor: EmptyProcessor,
  }
}

func FArray(a []Any) Monad {
  return Monad {
    content: a,
    processor: EmptyProcessor,
  }
}

func (m Monad) Filter(p Predicate) Monad {

  proc := func(e Any) Any {
    processed := m.processor(e)
    if processed == nil {
      return nil
    }

    if p(processed) {
      return processed
    } else {
      return nil // TODO return None
    }
  }

  out := Monad {
    content: m.content,
    processor: proc,
  }

  return out
}

func (m Monad) Map(f Functor) Monad {

  proc := func(e Any) Any {
    processed := m.processor(e)
    if processed == nil {
      return nil
    }
    return f(processed)
  }

  out := Monad {
    content: m.content,
    processor: proc,
  }

  return out
}

func (m Monad) ZipWithIndex() Monad {
  i := 0

  proc := func(e Any) Any {
    processed := m.processor(e)
    if processed == nil {
      return nil
    }
    currentIdx := i
    i = i + 1
    return Tuple2{currentIdx, processed}
  }

  out := Monad {
    content: m.content,
    processor: proc,
  }

  return out

}

func (m Monad) Foreach(f func(Any)) {
  for _, e := range m.content {
    processed := m.processor(e)
    if processed != nil {
      f(processed)
    }
  }
}


func (m Monad) ToChannel() chan Any {
  c := make(chan Any)

  go func() {
    defer close(c)

    m.Foreach(func(e Any) {
      c <- e
    })
  }()
 return c
}

func pos(e Any) bool {
  return e.(int) >= 0
}

func mult100(e Any) Any {
  return e.(int) * 100
}

func invertor(e Any) Any {
  return e.(int) * (-1)
}

func decorate(e Any) Any {
  return fmt.Sprintf("[%d]", e.(int))
}

func logger(e Any) {
  fmt.Println(e)
}

func even(e Any) bool {
  return e.(int) % 2 == 0
}

func main(){
  fmt.Println("main")


  FSeq(1,-2,3,-4,-5,6,-7,-8,-9,10, -20, -30, -40, 22, 32, 34).
    Filter(pos).
    Filter(even).
    Map(mult100).
    Map(invertor).
    Map(decorate).
    ZipWithIndex().
    Foreach(logger)
    //ToChannel()

    xs := List { 1, &List{2, &List {3, &Nil}} }
    xs.Println()

}
