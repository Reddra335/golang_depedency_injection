package simple

type Database struct {
	Name string
}

type PostgreSql Database
type MongoDB Database

func NewPostgreSql() *PostgreSql {
	return (*PostgreSql)(&Database{Name: "Postgre Sql"})
}
func NewMongoDB() *MongoDB {
	return (*MongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	*PostgreSql
	*MongoDB
}

func NewDatabaseRepository(postgreSql *PostgreSql, mongoDB *MongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		PostgreSql: postgreSql,
		MongoDB:    mongoDB,
	}
}
