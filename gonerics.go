package main

import (
	"fmt"
	"math/rand"
	"time"
)

type pile[T any] struct {
	items []T
}

func newPile[T any] () pile[T] {
	return pile[T]{
		items: nil,
	}
}

func (p *pile[T]) add(item T) {
	p.items = append(p.items, item)
}

func (p *pile[T]) get() *T {
	if len(p.items) == 0 {
		return nil
	}
	i := rand.Intn(len(p.items))
	result := p.items[i]
	p.items[i] = p.items[len(p.items)-1]
	p.items = p.items[:len(p.items)-1]
	return &result
}

type foo interface {
	fooo()
}
type fooOne struct {
}
func (f fooOne) fooo() {}
type fooTwo struct {
}
func (f fooTwo) fooo() {}
type bar interface {
	barr()
}
type barOne struct {
}
func (b barOne) barr() {}
type barTwo struct {
}
func (b barTwo) barr() {}

func match[T any] (a, b T) {}

func main() {
	rand.Seed(time.Now().UnixNano())
	ints := newPile[int]()
	for i := 0; i < 3; i++ {
		ints.add(i)
	}
	for i := 0; i < 4; i++ {
		result := ints.get()
		if result != nil {
			fmt.Printf("%v\n", *result)
		} else {
			fmt.Printf("done!\n")
		}
	}
	match[foo](fooOne{}, fooTwo{})
	// This won't compile
	// match[foo](fooOne{}, barTwo{})
	match[bar](barOne{}, barTwo{})
}
