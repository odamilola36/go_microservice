package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"microservices/authentication/models"
	"microservices/db"
)

const userCollection = "users"

type UserRepository interface {
	save(user *models.User) error
	getById(id string) (models.User, error)
	getByEmail(email string) (models.User, error)
	getAll() ([]models.User, error)
	updateUser(u *models.User) error
	deleteUser(id string) error
	deleteAll() error
}

type usersRepository struct {
	c *mgo.Collection
}

func NewUsersRepository(c db.Connection) UserRepository {
	return &usersRepository{
		c: c.DB().C(userCollection),
	}
}

func (r *usersRepository) save(user *models.User) error {
	return r.c.Insert(user)
}

func (r *usersRepository) getById(id string) (models.User, error) {
	var user models.User
	err := r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (r *usersRepository) getByEmail(email string) (models.User, error) {
	var user models.User
	err := r.c.Find(bson.M{"email": email}).One(&user)
	return user, err
}

func (r *usersRepository) getAll() ([]models.User, error) {
	var user []models.User
	err := r.c.FindId(bson.M{}).One(&user)
	return user, err
}

func (r *usersRepository) updateUser(u *models.User) error {
	var err = r.c.UpdateId(bson.M{"_id": u.Id}, u)
	return err
}

func (r *usersRepository) deleteUser(id string) error {
	var err = r.c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *usersRepository) deleteAll() error {
	err := r.c.DropCollection()
	return err
}
