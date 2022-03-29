package redis

type OprationError struct {
	opration string
}

func (err *OprationError) Error() string {
	return "Could not perform the " + err.opration + " operation."
}

type DownError struct{}

func (dbe *DownError) Error() string {
	return "Database is down."
}

type CreateDatabaseError struct{}

func (err *CreateDatabaseError) Error() string {
	return "Could not create Database"
}

type NoImplementedDatabaseError struct {
	database string
}

func (err *NoImplementedDatabaseError) Error() string {
	return err.database + " not implemented."
}
