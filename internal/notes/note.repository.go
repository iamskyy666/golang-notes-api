package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// repo -> data access layer
// all mongoDB queries here so that our handlers remain clean (talking to the DB here)
// CLEAN ARCHITECTURE


type Repo struct{
	coll *mongo.Collection // 1 single mongo collection, like a table
}

func NewRepo(db *mongo.Database)*Repo{
	return &Repo{
		coll: db.Collection("notes"),
	}
}

// DB Methods:
func (r *Repo)Create(ctx context.Context, note Note)(Note, error){
	opCtx,cancel:=context.WithTimeout(ctx, 5*time.Second) // opCtx -> child context
	defer cancel()

	_,err:=r.coll.InsertOne(opCtx,note)

	if err != nil {
		return Note{},fmt.Errorf("⚠️ Inserting Note Failed: %w",err)
	}
	return note,nil
}