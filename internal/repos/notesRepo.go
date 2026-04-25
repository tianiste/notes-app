package repos

import "go.mongodb.org/mongo-driver/v2/mongo"

type NotesRepo struct {
	Collection *mongo.Collection
}

func NewNotesRepo(db *mongo.Database) *NotesRepo {
	return &NotesRepo{Collection: db.Collection("notes")}
}
