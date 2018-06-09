package models

import (
	"log"

	"github.com/Electra-project/electra-auth/src/database"
	"github.com/globalsign/mgo/bson"
)

// User model.
type User struct {
	ID               string   `bson:"_id"`
	Challenge        string   `json:"challenge"`
	ChallengeAnswer  string   ``
	PurseHash        string   `json:"purseHash"`
	PursePrivateKey  string   ``
	TwitterID        string   `json:"twitterId"`
	TwitterCheckedAt int64    ``
	BootstrapNodes   []string `json:"bootstrapNodes"`
	CreatedAt        int64    `json:"createdAt"`
	UpdatedAt        int64    `json:"updatedAt"`
}

// GetByPurseHash gets a user from the database by their Purse Account address hash.
func (h User) GetByPurseHash(purseHash string) (*User, error) {
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
