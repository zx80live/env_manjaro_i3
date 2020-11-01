package fp

import "fmt"

type Any = interface{}
type Functor func (Any) Any
type Monad chan Any
type Predicate func (Any) bool
type Option = Monad
type Traversable = Monad
type Tuple2 struct {
  _1 Any
  _2 Any
}
type Memoized func(Any) Any

func Memoize(m Memoized) Memoized {
  cache := make(map[Any]Any)

  return func(key Any) Any {
    if val, found := cache[key]; found {
      return val
    }
    temp := m(key)
    cache[key] = temp
    return temp
  }
}

func Cons(elements ... Any) Monad {
  c := make(Monad)

  go func () {
    defer close(c)
    for _, e := range elements {
      c <- e
    }
  }()

  return c
}

func (m Monad) Foreach(fn func(Any)) {
  for e := range m {
    fn(e)
  }
}

func (m Monad) Map(fn Functor) Monad {
  c := make(Monad)

  go func() {
    defer close(c)
    for e := range m {
      c <- fn(e)
    }
  }()
  return c
}

func (m Monad) Filter(p Predicate) Monad {
  c := make(Monad)

  go func() {
    defer close(c)
    for e:= range m {
      if res := p(e); res {
        c <- e
      }
    }
  }()

  return c
}

func (m Monad) FilterNot(p Predicate) Monad {
  notP := func(e Any) bool {
    return !p(e)
  }
  return m.Filter(notP)
}

func (m Monad) Exists(p Predicate) bool {
  exist := false
  for e := range m {
    if p(e) {
      exist = true
      break
    }
  }

  return exist
}

func (m Monad) Count(p Predicate) int {
  count := 0

  for range m {
    count = count + 1
  }

  return count
}

func (m Monad) Forall(p Predicate) bool {
  count := 0
  trues := 0

  for e := range m {
    count = count + 1
    if p(e) {
      trues = trues + 1
    }
  }

  return trues > 0 && trues == count
}

func (m Monad) ZipWithIndex() Monad {
  c := make(Monad)

  go func(){
    defer close(c)
    count := 0
    for e := range m {
      c <- Tuple2{count, e}
      count = count + 1
    }

  }()

  return c
}

func (m Monad) Fold(init Any, fn func(Any, Any) Any ) Any {
  var acc Any
  acc = init

  for e := range m {
    acc = fn(e, acc)
  }
 
  return acc
}

func (m1 Monad) Zip(m2 Monad) Monad {
  fmt.Println("invoke Zip")

  c := make(Monad)

  guard := 30

  z1 := m1.ZipWithIndex()
  z2 := m2.ZipWithIndex()

  go func(){
    defer close(c)

    for {
      var t1 Tuple2
      var t2 Tuple2

      select {
      case e1:=<-z1:

        if e1 != nil {
          t1 = e1.(Tuple2)
          //fmt.Println("from m1: ", t1._1, t1._2)
          //c <- e1
        } else {
          //fmt.Println("m1 is exosted")
        }

      case e2:=<-z2:
        if e2 != nil {
          t2 = e2.(Tuple2)
          //fmt.Println("from m2: ", t2._1, t2._2)
          //c <- e2
        } else {
          //fmt.Println("m2 is exosted")
        }
      }

      //fmt.Println(" ******* ", t1, t2)
      c <- Tuple2 { t1, t2 }

      guard = guard - 1

      if guard <= 0 {
        break
      }
    }


  }()

  return c
}
