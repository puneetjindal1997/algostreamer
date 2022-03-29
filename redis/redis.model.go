package redis

type Database interface {
	Set(key string, value interface{}) (string, error)
	Get(key string) (interface{}, error)
	Delete(key string) (string, error)
	HSet(key string, field string, value interface{}) error
	HGet(key string, field string) (interface{}, error)
}

/*
 *	Function to create redis database and check if it is implemented or not
 */
func Factory(databaseName string) (Database, error) {
	switch databaseName {
	case "test":
		return createRedisDatabase()
	default:
		return nil, &NoImplementedDatabaseError{databaseName}
	}
}
