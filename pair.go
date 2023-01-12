package imc

type Pair[A any, B any] func() (A, B)

func MakePair[A any, B any](a A, b B) Pair[A, B] {
    return func() (A, B) {
        return a, b
    }
}
func (p Pair[A, B]) First() A {
    var a, _ = p()
    return a
}
func (p Pair[A, B]) Second() B {
    var _, b = p()
    return b
}
func (p Pair[K, V]) Key() K {
    return p.First()
}
func (p Pair[K, V]) Value() V {
    return p.Second()
}
