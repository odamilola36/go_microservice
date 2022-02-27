package dbCon

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection interface {
	Close(ctx context.Context) error
	DB() *mongo.Database
	DBContext() context.Context
}

type connection struct {
	database *mongo.Database
	c        context.Context
}

func NewConnection(c Config) (Connection, error) {
	serverOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(c.Dsn()).
		SetServerAPIOptions(serverOptions)
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	database := client.Database(c.DbName())
	return &connection{
		database: database,
		c:        nil,
	}, nil
}

func (c connection) Close(ctx context.Context) error {
	return c.database.Client().Disconnect(ctx)
}

func (c connection) DBContext() context.Context {
	return c.c
}

func (c connection) DB() *mongo.Database {
	return c.database
}
