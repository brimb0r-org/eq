package component_tests

import (
	"context"
	"fmt"
	"gSheets/application/internal/eq_repo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
	"time"
)

var (
	numberOfRecords = 1000000
)

func TestEndToEnd(t *testing.T) {
	if ComponetTests == "" {
		t.Skip()
	}
	configuration := configureTests()
	database := configuration.Mongo.MongoDatabase()
	populateDatabase(database, t)
	start := time.Now()
	eqRepo := &eq_repo.Repo{Database: database}
	_, err := eqRepo.QueryEq(context.Background())
	assert.NoError(t, err)

	dbDuration := time.Since(start)
	log.Printf("Query took Time [%s] to pull [%d] records", dbDuration, numberOfRecords)

}

func populateDatabase(repo *mongo.Database, t *testing.T) {
	ctx := context.Background()
	collection := repo.Collection("eq")

	err := collection.Drop(ctx)
	assert.NoError(t, err)

	eq := make([]interface{}, numberOfRecords)
	for i := 0; i < numberOfRecords; i++ {
		eq[i] = eq_repo.Eq{
			ID:   fmt.Sprintf("%d", i),
			Name: "eq",
		}
	}
	_, err = collection.InsertMany(ctx, eq)
	assert.NoError(t, err)
}
