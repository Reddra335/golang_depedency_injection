package simple

type FooBar struct {
	*BarService
	*FooService
}

func NewFooBar(bar *BarService, foo *FooService) *FooBar {
	return &FooBar{
		BarService: bar,
		FooService: foo,
	}
}
