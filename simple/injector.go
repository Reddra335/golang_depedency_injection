//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializedService(iSError bool) (*SimpleService, error) {

	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabase() *DatabaseRepository {
	wire.Build(
		NewPostgreSql,
		NewMongoDB,
		NewDatabaseRepository,
	)
	return nil
}

var setFoo = wire.NewSet(NewFooRepository, NewFooService)
var setBar = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBar() *FooBar {
	wire.Build(setBar, setFoo, NewFooBar)
	return nil
}

//salah
// func InitializedHelloService() *HelloService {
// 	wire.Build(NewSayHelloImpln, NewHelloService)
// 	return nil
// }

var setHello = wire.NewSet(NewSayHelloImpln, wire.Bind(new(SayHello), new(*SayHelloImpln)))

func InitializedHelloService() *HelloService {

	wire.Build(setHello, NewHelloService)

	return nil
}

func InitializedFoo_Bar() *Foo_Bar {

	wire.Build(NewBar, NewFoo, wire.Struct(new(Foo_Bar), "Bar", "Foo"))

	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *Foo_Bar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(Foo_Bar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(NewAplication, wire.FieldsOf(new(*Aplication), "Configuration"))
	return nil
}
