package fp

type List struct {
  head Any
  tail *List
  functor Functor
}

var Nil List = List { nil, nil, EmptyFunctor }

func (l List) Head() Any {
  return l.head
}

func (l List) Tail() List {
  return l.tail.Copy()
}

func (l List) IsEmpty() bool {
  return l.head == nil && l.tail == nil
}

func (l List) IsNotEmpty() bool {
  return !l.IsEmpty()
}

/**
 * Copy (references only) list with empty functors
 */
func (l List) Copy() List {
  if l.IsEmpty() {
    return l
  } else {
    tail := l.tail.Copy()
    return List {
      head : l.head,
      tail : &tail,
      functor : EmptyFunctor,
    }
  }
}

/**
 * Add head to list
 *
 * O(1)
 */
func (l List) Cons(e Any) List {
  tail := l.Copy()
  xs := List {
    head: e,
    tail: &tail,
    functor: EmptyFunctor,
  }

  return xs
}

/**
 * Add lazy filter
 */
func (l List) Filter(p Predicate) List {
  return List {
    head : l.head,
    tail : l.tail,
    functor : func(e Any) Any {
      if processed := l.functor(e); processed != nil && p(processed) {
          return processed
      } else {
          return nil                                    //TODO return None
      }
    },
  }
}

/**
 * Add lazy functor
 */
func (l List) Map(f Functor) List {
  return List {
    head : l.head,
    tail : l.tail,
    functor : func(e Any) Any {
      if processed := l.functor(e); processed != nil {
        return f(processed)
      } else {
        return nil
      }
    },
  }
}

func (l List) Reverse() List {
  xs := Nil
  l.Foreach(func(e Any) {
    xs = xs.Cons(e)
  })
  return xs
}

func (l1 List) Zip(l2 List) List {
  zipped := Nil
  it1 := &l1
  it2 := &l2
  for {
    if (it1.head != nil && it2.head != nil) {
      zipped = zipped.Cons(Tuple2 { it1.head, it2.head })
      it1 = it1.tail
      it2 = it2.tail
    } else {
      break
    }
  }

  return zipped.Reverse()
}

func (l List) ZipWithIndex() List {
  zipped := Nil
  i := 0
  it := &l

  for {
    if (it.head != nil) {
      zipped = zipped.Cons(Tuple2 { it.head, i})
      it = it.tail
      i = i + 1
    } else {
      break
    }
  }

  return zipped.Reverse()
}

func (l List) Count(p Predicate) int {
  c := 0
  l.Foreach(func(e Any) {
    if p(e) {
      c = c + 1
    }
  })
  return c
}

func (l List) Exist(p Predicate) bool {
  e := false
  it := &l
  for {
    if it.head != nil {
      if p(it.head) {
        e = true
        break
      } else {
        it = it.tail
      }
    } else {
      break
    }
  }

  return e
}

//TODO optimize: get rid of head recursion
func (l List) Reduce(f func(el Any, er Any) Any) Any {
  if l.IsEmpty() {
    panic("Usupport operation: reduce left on empty list")
  } else if l.tail.IsEmpty() {
    return l.head
  } else {
    return f(l.head, l.tail.Reduce(f))
  }
}


/**
 * = O(n)
 */
func (l List) Size() int {
  return l.Count(func(e Any) bool {return true})
}

/**
 * Materialize list: apply functors, filters before each element will be passed to the lambda
 */
func (l List) Foreach(f func(Any)) {
  if l.head != nil {
    processed := l.functor(l.head)
    if processed != nil {
      f(processed)
    }
  }

  if l.tail != nil {
    l.tail.mapHead(l.functor).Foreach(f)
  }
}


//-- internal --------------------------------------------------------------------

/**
 * Add lazy functor to the head only
 */
func (l List) mapHead(f Functor) List {
  return List {
    head: l.head,
    tail: l.tail,
    functor : func(e Any) Any {
      processed := l.functor(e)
      if processed == nil {
        return nil
      } else {
        return f(processed)
      }
    },
  }
}

