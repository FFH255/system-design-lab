package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		collection: client.Database("test").Collection("groups"),
	}
}

func (r *Repository) Get(ctx context.Context, uid UID) (*Group, error) {
	filter := bson.M{"uid": uid}

	var group Group

	err := r.collection.FindOne(ctx, filter).Decode(&group)

	if err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return &group, nil
}

func (r *Repository) Save(ctx context.Context, group *Group) error {
	filter := bson.M{"uid": group.UID}
	update := bson.M{"$set": group}
	opts := options.Update().SetUpsert(true) // Creates a new document if not found

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, uid UID) error {
	filter := bson.M{"uid": uid}

	_, err := r.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
