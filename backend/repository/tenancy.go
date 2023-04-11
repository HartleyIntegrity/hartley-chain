package repository

import (
	"hartley-chain/backend/models"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database
var collection *mgo.Collection

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB("property-blockchain-system")
	collection = db.C("tenancies")
}

func GetNewID() bson.ObjectId {
	return bson.NewObjectId()
}

func GetAllTenancies() ([]models.Tenancy, error) {
	var tenancies []models.Tenancy
	err := collection.Find(nil).All(&tenancies)
	return tenancies, err
}

func GetTenancy(id string) (models.Tenancy, error) {
	var tenancy models.Tenancy
	err := collection.FindId(bson.ObjectIdHex(id)).One(&tenancy)
	return tenancy, err
}

func AddTenancy(tenancy models.Tenancy) error {
	err := collection.Insert(&tenancy)
	return err
}

func UpdateTenancy(tenancy models.Tenancy) error {
	err := collection.UpdateId(tenancy.ID, tenancy)
	return err
}

func DeleteTenancy(id string) error {
	err := collection.RemoveId(bson.ObjectIdHex(id))
	return err
}
