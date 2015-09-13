package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Game struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Field Field
}

var session *mgo.Session = nil

func GetSession() *mgo.Session {
	if session == nil {
		new_session, err := mgo.Dial("127.0.0.1:27017")
		if err != nil {
			panic(err)
		}
		session = new_session
	}
	return session
}

func MakeDBNewGame(collection *mgo.Collection) Game {
	game := Game{bson.NewObjectId(), InitGameField()}
	err := collection.Insert(&game)
	if err != nil {
		log.Fatal(err)
	}
	return game
}

func GetGameById(id bson.ObjectId, collection *mgo.Collection) *Game {
	game := Game{Id: id, Field: Field{}}
	err := collection.FindId(game.Id).One(&game)
	if err != nil {
		log.Fatal(err)
	}
	return &game
}

func SaveGameById(game *Game, collection *mgo.Collection) {
	if _, err := collection.UpsertId(game.Id, game); err != nil {
		log.Fatal(err)
	}
}
