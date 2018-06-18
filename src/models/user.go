package models

import (
	"log"
	"time"

	"github.com/Electra-project/electra-auth/src/database"
	"github.com/Electra-project/electra-auth/src/helpers"
	"github.com/globalsign/mgo/bson"
)

// User model.
type User struct {
	ID               bson.ObjectId   `bson:"_id" json:"-"`
	PurseHash        string          `bson:"purseHash" json:"purseHash"`
	Token            string          `bson:"token" json:"token"`
	PursePrivateKey  string          `bson:"pursePrivateKey" json:"-"`
	TwitterUsername  string          `bson:"twitterUsername" json:"twitterUsername"`
	TwitterCheckedAt time.Time       `bson:"twitterCheckedAt" json:"twitterCheckedAt"`
	BootstrapNodes   []bson.ObjectId `bson:"bootstrapNodes" json:"-"`
	CreatedAt        time.Time       `bson:"createdAt" json:"createdAt"`
	UpdatedAt        time.Time       `bson:"updatedAt" json:"updatedAt"`
}

const tokenLength uint8 = 196

// GetByPurseHash gets a user from the database
// by their Purse Account address hash.
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
	token, err := helpers.GenerateToken(tokenLength)
	if err != nil {
		return nil, err
	}

	db := database.Get()
	collection := db.C("users")

	err = collection.Insert(bson.M{
		"purseHash": purseHash,
		"token":     token,
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return u.GetByPurseHash(purseHash)
}
