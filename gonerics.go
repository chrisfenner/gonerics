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

type producer[T any] interface {
	produce() T
}

func (b barOne) produce() barTwo {
	return barTwo{}
}

func (b fooOne) produce() fooTwo {
	return fooTwo{}
}

func produce[B any](a producer[B]) B {
	return a.produce()
}

func produce2[B any](a producer[B], b *B) {
	*b = a.produce()
}

func produce3[B any, A producer[B]](a A, b *B) {
	*b = a.produce()
}

type concreteProducer[T any] struct {
}

type concreteFooOne struct {
	concreteProducer[fooTwo]
}

func (c concreteProducer[T]) produce() T {
	return *new(T)
}

func produce4[B any](a concreteProducer[B]) B {
	return a.produce()
}

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
	match[bar](barOne{}, barTwo{})
	// This won't compile, it produces the error:
	// cannot use barTwo{} (value of type barTwo) as type foo in
	// argument to match[foo]:
	// barTwo does not implement foo (missing fooo method)
	// match[foo](fooOne{}, barTwo{})

	// This won't compile either:
	// type fooTwo of fooTwo{} does not match inferred type
	// fooOne for T
	// match(fooOne{}, fooTwo{})

	f := fooOne{}
	var f2 fooTwo

	// This won't compile, because the compiler only performs type
	// inference for parameters (not return values).
	// f2 = produce(f)
	// Instead, we still have to explicitly provide the type:
	f2 = produce[fooTwo](f)

	// We can fix this by using an output pointer parameter to
	// trick the caller into giving us the type.

	// This won't compile, because the compiler still can't infer
	// fooTwo just based on parameter type inference.
	// produce2(f, &f2)
	// Instead, we still have to explicitly provide the type:
	produce2[fooTwo](f, &f2)

	// Now the compiler can infer the correct types [fooOne, fooTwo]
	// to use with produce3.
	produce3(f, &f2)

	// We can embed a type with type parameters and use it with
	// produce3
	cf := concreteFooOne{}
	produce3(cf, &f2)

	// We can also (obviously) just call the method on that
	f2 = cf.produce()

	// We still can't use the produce function with it.
	// The type inferencer still can't infer B.
	// f2 = produce(cf)
	// So we can explicitly provide the type.
	f2 = produce[fooTwo](cf)
}
