package repositories

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"microservices/authentication/models"
	"microservices/db"
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
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.save(user)
	assert.NoError(t, err)

	found, err := r.getById(id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)

}

func TestUserRepositoryGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.save(user)
	assert.NoError(t, err)

	found, err := r.getById(id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)

	found, err = r.getById(bson.NewObjectId().Hex())
	assert.Nil(t, found)
	assert.Error(t, mgo.ErrNotFound, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())

}

func TestUserRepositoryGetByEmail(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.save(user)
	assert.NoError(t, err)

	found, err := r.getByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user, found)

	found, err = r.getById("")
	assert.Nil(t, found)
	assert.Error(t, mgo.ErrNotFound, err)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
}

func TestUserRepositoryUpdate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.save(user)
	assert.NoError(t, err)

	user.Name = "Test2"

	err = r.updateUser(user)
	assert.NoError(t, err)

	found, err := r.getById(id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, found.Name, user.Name)

}

func TestUserRepositoryDelete(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.save(user)
	assert.NoError(t, err)

	err = r.deleteUser(user.Id.Hex())
	assert.NoError(t, err)

	found, err := r.getById(id.Hex())
	assert.Error(t, mgo.ErrNotFound, err)
	assert.Nil(t, found)
	assert.EqualError(t, mgo.ErrNotFound, err.Error())
}

func TestUserRepositoryDeleteAll(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id:       id,
		Name:     "Test",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "123456789",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.save(user)
	assert.NoError(t, err)

	names, err := conn.DB().CollectionNames()
	assert.NoError(t, err)
	b := len(names) > 0
	assert.Equal(t, b, true)

	err = r.deleteAll()
	assert.NoError(t, err)
	b = len(names) == 0
	assert.Equal(t, b, true)
}
