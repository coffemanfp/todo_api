package database

// Database is the Database manager for connections and repository instancies.
type Database struct {
	Conn         DatabaseConnector
	Repositories Repositories
}

// DatabaseConnector defines a database connector handler.
type DatabaseConnector interface {

	// Connect creates new connection of the database implementation.
	//  @return $1 error: database connection error
	Connect() error
}
