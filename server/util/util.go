package util

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckId(id string) error {
	if id == "" || len(id) != 24 {
		return errors.New(ParameterError)
	}
	return nil
}

func ToObjectId(id string) primitive.ObjectID {
	if id != "" && len(id) == 24 {
		hex, _ := primitive.ObjectIDFromHex(id)
		return hex
	}
	return primitive.NilObjectID
}
