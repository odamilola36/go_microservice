package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type Connection interface {
	Close()
	DB() *mgo.Database
}

//mongodb+srv://<username>:<pas
//sword>@ms-cluster.3vorc.mongodb.net/test
//lomari
//mK5wqaKBYTtENjq5

type connection struct {
	session  *mgo.Session
	database *mgo.Database
}

func NewConnection(c Config) (Connection, error) {
	fmt.Printf("Connecting to mongo %v \n", c.Dsn())
	session, err := mgo.Dial(c.Dsn())

	if err != nil {
		return nil, err
	}

	return &connection{session, session.DB(c.DbName())}, nil
}

func (c *connection) Close() {
	c.session.Close()
}

func (c *connection) DB() *mgo.Database {
	return c.database
}
