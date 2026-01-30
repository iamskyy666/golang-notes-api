package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(r *gin.Engine, db *mongo.Database){
	// create repo and handler once, at STARTUP
	repo:=NewRepo(db)
	h:=NewHandler(repo)

	notesGroup:=r.Group("/notes")
	{
		notesGroup.POST("",h.CreateNote)
		notesGroup.GET("",h.ListNotes)
	}
}