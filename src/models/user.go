package models

import (
	"log"
	"time"

	"github.com/Electra-project/electra-auth/src/database"
	"github.com/globalsign/mgo/bson"
)

// User model.
type User struct {
	ID               string   `bson:"_id"`
	PurseHash        string   `json:"purseHash"`
	PursePrivateKey  string   ``
	TwitterUsername  string   `json:"twitterUsername"`
	TwitterCheckedAt int64    ``
	BootstrapNodes   []string `json:"bootstrapNodes"`
	CreatedAt        int64    `json:"createdAt"`
	UpdatedAt        int64    `json:"updatedAt"`
}

// GetByPurseHash gets a user from the database by their Purse Account address hash.
func (u User) GetByPurseHash(purseHash string) (*User, error) {
	db := database.Get()
	collection := db.C("users")

	var user *User
	err := collection.Find(bson.M{"purseHash": purseHash}).One(&user)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return user, nil
}

// Insert creates a new user in the database.
func (u User) Insert(purseHash string) (*User, error) {
	db := database.Get()
	collection := db.C("users")

	err := collection.Insert(bson.M{
		"purseHash":        purseHash,
		"PursePrivateKey":  nil,
		"twitterUsername":  nil,
		"twitterCheckedAt": nil,
		"bootstrapNodes":   [0]string{},
		"createdAt":        time.Now(),
		"updatedAt":        time.Now(),
	})
	if err != nil {
		return nil, err
	}

	var user *User
	err = collection.Find(bson.M{"purseHash": purseHash}).One(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
