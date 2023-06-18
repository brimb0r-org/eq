package eq_repo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Eq struct {
	ID             string `bson:"_id"`
	Name           string
	Activity       string
	Publish_status bool
	LastUpdated    primitive.DateTime `bson:"last_updated,omitempty"`
}
