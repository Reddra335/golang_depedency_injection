//go:build wireinject
// +build wireinject

package simple

import (
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
