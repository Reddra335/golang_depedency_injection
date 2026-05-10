package simple

type Foo struct{}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct{}

func NewBar() *Bar {
	return &Bar{}
}

type Foo_Bar struct {
	*Foo
	*Bar
}

func NewFoo_Bar(foo *Foo, bar *Bar) *Foo_Bar {
	return &Foo_Bar{
		Foo: foo,
		Bar: bar,
	}
}
