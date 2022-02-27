package repositories

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"microservices/authentication/models"
	"microservices/dbCon"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
}

func TestUserRepositorySave(t *testing.T) {
	cfg := dbCon.NewConfig()
	conn, err := dbCon.NewConnection(cfg)
	assert.NoError(t, err)
	defer func(conn dbCon.Connection, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, conn.DBContext())

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)

}

func TestUserRepositoryGetById(t *testing.T) {
	cfg := dbCon.NewConfig()
	conn, err := dbCon.NewConnection(cfg)
	assert.NoError(t, err)
	defer func(conn dbCon.Connection, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, conn.DBContext())

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)

	found, err = r.GetById(primitive.NewObjectID().Hex())
	assert.Equal(t, found.Name, "")
	assert.Error(t, mongo.ErrNilValue, err)
	assert.EqualError(t, mongo.ErrNoDocuments, err.Error())

}

func TestUserRepositoryGetByEmail(t *testing.T) {
	cfg := dbCon.NewConfig()
	conn, err := dbCon.NewConnection(cfg)
	assert.NoError(t, err)
	defer func(conn dbCon.Connection, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, conn.DBContext())

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	//assert.Equal(t, user, found)

	found, err = r.GetById("")
	assert.Equal(t, found.Email, "")
	assert.Error(t, mongo.ErrNilValue, err)
	assert.EqualError(t, mongo.ErrNoDocuments, err.Error())
}

func TestUserRepositoryUpdate(t *testing.T) {
	cfg := dbCon.NewConfig()
	conn, err := dbCon.NewConnection(cfg)
	assert.NoError(t, err)
	defer func(conn dbCon.Connection, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, conn.DBContext())

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	user.Name = "Test2"

	err = r.UpdateUser(user)
	assert.NoError(t, err)

	found, err := r.GetById(id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, found.Name, user.Name)

}

func TestUserRepositoryDelete(t *testing.T) {
	cfg := dbCon.NewConfig()
	conn, err := dbCon.NewConnection(cfg)
	assert.NoError(t, err)
	defer func(conn dbCon.Connection, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, conn.DBContext())

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	err = r.DeleteUser(user.Id.Hex())
	assert.NoError(t, err)

	found, err := r.GetById(id.Hex())
	assert.Error(t, mongo.ErrNilDocument, err)
	assert.Equal(t, found.Email, "")
	assert.EqualError(t, mongo.ErrNoDocuments, err.Error())
}

func TestUserRepositoryDeleteAll(t *testing.T) {
	cfg := dbCon.NewConfig()
	conn, err := dbCon.NewConnection(cfg)
	assert.NoError(t, err)
	defer func(conn dbCon.Connection, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, conn.DBContext())

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	//names, err := conn.DB().CollectionNames()

	names, err := conn.DB().ListCollectionNames(conn.DBContext(), bson.M{})
	assert.NoError(t, err)
	b := len(names) > 0
	assert.Equal(t, b, true)

	err = r.DeleteAll()
	assert.NoError(t, err)
	names, err = conn.DB().ListCollectionNames(conn.DBContext(), bson.M{})
	b = len(names) == 0
	assert.Equal(t, b, true)
}
