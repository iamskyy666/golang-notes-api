package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// DB CRUD Methods:

// Insert-> Create
func (r *Repo)Create(ctx context.Context, note Note)(Note, error){
	opCtx,cancel:=context.WithTimeout(ctx, 5*time.Second) // opCtx -> child context
	defer cancel()

	_,err:=r.coll.InsertOne(opCtx,note)

	if err != nil {
		return Note{},fmt.Errorf("⚠️ Inserting Note Failed: %w",err)
	}
	return note,nil
}

// List-> Read
func (r *Repo)List(ctx context.Context)([]Note, error){
	opCtx,cancel:=context.WithTimeout(ctx, 5*time.Second) // opCtx -> child context
	defer cancel()

	//bson.M{} - used to build filter{}, empty {} - all docs.
	filter:=bson.M{}

	// cursor - kinda iterator
	cursor,err:=r.coll.Find(ctx,filter)

	if err != nil {
		return nil, fmt.Errorf("⚠️ Failed to find notes: %w",err)
	}

	defer cursor.Close(opCtx) // cursor must be closed after use. (to avoid leaks, as it uses up server-resources)

	var notes []Note

	//cursor.All() -> read all the remaining docs. from the particular cursor into a slice
	if err:=cursor.All(opCtx, &notes);err!=nil{
		return nil, fmt.Errorf("⚠️ Failed to decode notes: %w",err)
	} 

	return notes,nil
}

// FindById-> Read
func (r *Repo)FindNote(ctx context.Context, id primitive.ObjectID)(Note, error){
	opCtx,cancel:=context.WithTimeout(ctx, 5*time.Second) // opCtx -> child context
	defer cancel()

	filter:=bson.M{"_id":id} // check model

	var note Note

	err:=r.coll.FindOne(opCtx,filter, options.FindOne()).Decode(&note)
	if err != nil {
		return Note{}, fmt.Errorf("⚠️ Failed to fetch note by ID: %w",err)
	}

	return  note,nil
}

