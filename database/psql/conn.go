package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type properties struct {
	user string
	pass string
	name string
	host string
	port int
}

// PostgreSQLConnector implements a database.DatabaseConnector handler.
//
//	It is a handler for the PostgreSQL connections.
type PostgreSQLConnector struct {
	props properties
	db    *sql.DB
}

func (p *PostgreSQLConnector) Connect() (err error) {
	db, err := sql.Open("postgres", connURL(p.props))
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("failed to ping database: %s", err)
		return
	}
	p.db = db
	return
}

func (p PostgreSQLConnector) getConn() (conn *sql.DB, err error) {
	if p.db == nil {
		err = p.Connect()
		if err != nil {
			return
		}
	}
	err = p.db.Ping()
	if err != nil {
		err = fmt.Errorf("failed to ping database: %s", err)
		return
	}

	conn = p.db
	return
}

// NewPostgreSQLConnector initializes a new *PostgreSQLConnector.
//
//	@param user string: PostgreSQL connection property.
//	@param pass string: PostgreSQL connection property.
//	@param name string: PostgreSQL connection property.
//	@param host string: PostgreSQL connection property.
//	@param port int: PostgreSQL connection property.
//	@return conn *PostgreSQLConnector: new *PostgreSQLConnector instance.
func NewPostgreSQLConnector(user, pass, name, host string, port int) (conn *PostgreSQLConnector) {
	return &PostgreSQLConnector{
		props: properties{
			user: user,
			pass: pass,
			name: name,
			host: host,
			port: port,
		},
	}
}

func connURL(props properties) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", props.user, props.pass, props.host, props.port, props.name, "disable")
}
