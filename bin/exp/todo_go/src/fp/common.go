package fp

type Any interface{}
type LazyAny func() Any //TODO implement
type Predicate func(e Any) bool
type Functor func(e Any) Any

var EmptyPredicate Predicate = func(e Any) bool { return true }
var EmptyFunctor Functor = func(e Any) Any { return e }

type Tuple2 struct {
  _1 Any
  _2 Any
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
