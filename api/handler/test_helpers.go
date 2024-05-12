package api

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	db.Store
}

func (tbd *testdb) teardown(t *testing.T) {
	if err := tbd.User.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return &testdb{
		Store: db.Store{
			User: db.NewMongoUserStore(client),
		},
	}
}
