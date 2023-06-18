package eq_repo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

const EqCollection = "eq"

type IEqRepo interface {
	QueryEq(ctx context.Context) ([]*Eq, error)
	UpdateEqPublished(eq *Eq) error
}

type Repo struct {
	*mongo.Database
}

func (r *Repo) QueryEq(ctx context.Context) ([]*Eq, error) {
	collection := r.Collection(EqCollection)
	results := make([]*Eq, 0)

	filter := bson.M{
		"Name":           "Eq",
		"Publish_status": false,
	}

	opts := options.Find().SetSort(bson.D{{"_id", 1}})

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = cursor.Close(ctx)
	}()

	for {
		if cursor.Err() != nil {
			return results, err
		}
		if ok := cursor.Next(ctx); !ok {
			break
		}

		eq := &Eq{}
		if err := cursor.Decode(eq); err != nil {
			return nil, err
		}
		results = append(results, eq)
	}

	return results, err
}

func (r *Repo) UpdateEqPublished(eq *Eq) error {
	_, err := r.Collection(EqCollection).UpdateOne(
		context.Background(),
		bson.D{{"_id", eq.ID}},
		bson.D{{"$set",
			bson.D{{"Publish_status", true}},
		}},
	)
	if err != nil {
		return fmt.Errorf("[%s] not updated in mongo: %w", eq.ID, err)
	}
	return nil
}
