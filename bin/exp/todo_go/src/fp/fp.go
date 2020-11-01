package fp

type Any = interface{}
type Functor func (Any) Any
type Monad chan Any
type Predicate func (Any) bool
type Option = Monad
type Traversable = Monad

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
