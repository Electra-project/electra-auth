package models

import (
	"log"
	"time"

	"github.com/Electra-project/electra-auth/src/database"
	"github.com/Electra-project/electra-auth/src/libs/mnemonic"
	"github.com/globalsign/mgo/bson"
)

// UserToken model.
type UserToken struct {
	ID        string    `bson:"_id" json:"-"`
	Challenge string    `json:"challenge"`
	PurseHash string    `bson:"purseHash" json:"purseHash"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// GetByPurseHash get or generate a user token from the database by their Purse
// Account address hash.
func (h UserToken) GetByPurseHash(purseHash string) (*UserToken, error) {
	db := database.Get()
	collection := db.C("users-tokens")

	count, err := collection.Find(bson.M{"purseHash": purseHash}).Count()
	if err != nil {
		log.Println(err)

		return nil, err
	}

	if count == 1 {
		var userToken *UserToken
		err := collection.Find(bson.M{"purseHash": purseHash}).One(&userToken)
		if err != nil {
			log.Println(err)

			return nil, err
		}

		return userToken, nil
	}

	challenge, err := generateMnemonic()
	if err != nil {
		log.Println(err)

		return nil, err
	}

	err = collection.Insert(bson.M{
		"challenge": challenge,
		"purseHash": purseHash,
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	})
	if err != nil {
		log.Println(err)

		return nil, err
	}

	var userToken *UserToken
	err = collection.Find(bson.M{"purseHash": purseHash}).One(&userToken)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return userToken, nil
}

func generateMnemonic() (string, error) {
	entropy, err := mnemonic.NewEntropy(256)
	if err != nil {
		log.Println(err)

		return "", err
	}

	mnemonic, err := mnemonic.NewMnemonic(entropy)
	if err != nil {
		log.Println(err)

		return "", err
	}

	return mnemonic, nil
}
