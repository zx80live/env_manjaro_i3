package fp

import (
	"fmt"
	"time"
)

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


func Zip22(m Monad) {
  for {
    fmt.Println("waiting monad elements")
    c := <-m
    fmt.Println("get element", c)
  }
}

func (m1 Monad) ZipWith(m2 Monad) Monad {
  fmt.Println("invoke ZipWith")
  time.Sleep(1 * time.Millisecond)
  
  c := make(Monad)

  go func(){
    defer close(c)
    for {
      el := <-m1
      c <- el
    }
  }()

  fmt.Println("exit from ZipWith")
  return c
}

//TODO Option
//TODO use lazy collections
//TODO fix order
func (m1 Monad) Zip(m2 Monad) Monad {
  fmt.Println("invoke Zip")

  c := make(Monad)
  tmp := make(Monad)


  go func(){
    defer close(tmp)
    defer close(c)

    guard := 100000

    z1 := m1.ZipWithIndex()
    z2 := m2.ZipWithIndex()

    for {
      var t1 Tuple2
      var t2 Tuple2

      select {
      case e1:=<-z1:
        if e1 != nil {
          t1 = e1.(Tuple2)
          tmp <- t1
        }
      case e2:=<-z2:
        if e2 != nil {
          t2 = e2.(Tuple2)
          tmp <- t2
        }
      }

      guard = guard - 1

      if guard <= 0 {
        break
      }
    }


  }()

  go func() {
    tmp.Fold(make(map[int]Any), func(e Any, acc Any) Any {
      el := e.(Tuple2)
      key := el._1.(int)
      cache := acc.(map[int]Any)

      if cached, found := cache[key]; found {
        c <- Tuple2 { cached, el._2 }
        delete(cache, key)
      } else {
        cache[key] = el._2
      }

      return cache
    })
  }()

  return c
}
