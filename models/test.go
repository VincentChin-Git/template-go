package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type test struct {
	Id        *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`             /* _id */
	CreatedAt *time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"` /* createdAt */
	UpdatedAt *time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"` /* updatedAt */
}
