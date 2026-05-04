package repos

import (
	"context"
	"notes-app/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type NotesRepo struct {
	Collection *mongo.Collection
}

func NewNotesRepo(db *mongo.Database) *NotesRepo {
	return &NotesRepo{Collection: db.Collection("notes")}
}

func (repo *NotesRepo) Create(ctx context.Context, note *models.Note) error {
	now := time.Now().UTC()

	note.ID = bson.NewObjectID()
	note.CreatedAt = now
	note.UpdatedAt = now

	_, err := repo.Collection.InsertOne(ctx, note)
	return err
}
