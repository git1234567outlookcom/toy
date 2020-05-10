package util

import (
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestProcessMongoErr(t *testing.T) {
	InitErrMap()
	println(ProcessMongoErr(mongo.ErrNoDocuments.Error()).Error())
}
