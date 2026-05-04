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

func (repo *NotesRepo) FindAll(ctx context.Context) ([]models.Note, error) {
	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notes []models.Note
	if err := cursor.All(ctx, &notes); err != nil {
		return nil, err
	}

	return notes, nil
}

func (repo *NotesRepo) FindByID(ctx context.Context, id string) (*models.Note, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var note models.Note
	err = repo.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&note)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (repo *NotesRepo) Update(ctx context.Context, id string, note models.Note) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	note.UpdatedAt = time.Now().UTC()
	update := bson.M{
		"$set": bson.M{
			"title":      note.Title,
			"body":       note.Body,
			"updated_at": note.UpdatedAt,
		},
	}

	_, err = repo.Collection.UpdateByID(ctx, objectID, update)
	return err
}

func (repo *NotesRepo) Delete(ctx context.Context, id string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = repo.Collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
