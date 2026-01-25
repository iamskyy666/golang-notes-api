package db

import (
	"context"
	"fmt"
	"time"

	"github.com/callmeskyy111/notes-api/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(cfg config.Config)(*mongo.Client, *mongo.Database,error){
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second * 10) // prevents app from freezing
	defer cancel()

	clientOpts:=options.Client().ApplyURI(cfg.MongoURI)

	client,err:=mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil,nil, fmt.Errorf("⚠️ FAILED, connecting MONGODB!")
	}

	// ping method
	if err:=client.Ping(ctx,nil);err!=nil{
		return nil,nil,fmt.Errorf("⚠️ Mongo-ping failed!")
	}

	database:=client.Database(cfg.MongoDB)

	return client,database,nil
}

func DisconnectDB(client *mongo.Client)error{
	ctx,cancel:=context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}