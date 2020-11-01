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
