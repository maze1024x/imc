package imc

import (
    "testing"
)

func TestMapBasic(t *testing.T) {
    var m = MakeMutMap[string, int](StringCompare,
        MakePair("a", 1),
        MakePair("b", 2),
    )
    m.Insert("c", 3)
    m.Insert("c", 4)
    m.Delete("b")
    {
        var v, found = m.Lookup("a")
        if !(found && v == 1) {
            t.Fatalf("wrong behavior for a")
        }
    }
    if m.Has("b") {
        t.Fatalf("wrong behavior for b")
    }
    {
        var v, found = m.Lookup("c")
        if !(found && v == 4) {
            t.Fatalf("wrong behavior for c")
        }
    }
    m.ForEach(func(key string, value int) {
        t.Log(key, value)
    })
}

// TODO: more test cases
