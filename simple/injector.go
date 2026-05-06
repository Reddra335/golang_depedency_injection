//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

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
