package notes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

// CRUD

func (h *Handler) CreateNote(ctx *gin.Context){
	var req CreateNoteRequest

	if err:= ctx.ShouldBindJSON(&req);err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"⚠️ Invalid JSON!",
			"status_code":http.StatusBadRequest,
		})
		return
	}

	now:=time.Now().UTC()

	note:=Note{
		ID: primitive.NewObjectID(),
		Title: req.Title,
		Content: req.Content,
		Pinned: req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	createdNote,err:=h.repo.Create(ctx.Request.Context(),note)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":"⚠️ Failed to CREATE note!",
			"status_code":http.StatusInternalServerError,
		})
		return
	}

	// If everything is OK.. ✅
	ctx.JSON(http.StatusCreated,createdNote)
}


func (h *Handler) ListNotes(ctx *gin.Context){

	notes,err:= h.repo.List(ctx.Request.Context())

	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":"⚠️ Failed to FETCH all notes!",
			"status_code":http.StatusInternalServerError,
		})
		return
	}

	// If everything is OK.. ✅
	ctx.JSON(http.StatusOK,gin.H{
		"message":"✅ Fetched all notes!",
		"notes":notes,
	})
}

func (h *Handler) GetNoteById(ctx *gin.Context){
	idStr:=ctx.Param("id")

	// ObjectIDFromHex creates a new ObjectID from a 24-char hex string. It returns an error if the hex string is not a valid ObjectID.
	objId,err:=primitive.ObjectIDFromHex(idStr)

	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"⚠️ Invalid ID!",
			"status_code":http.StatusBadRequest,
		})
		return
	}

	note,err:=h.repo.FindNote(ctx.Request.Context(),objId)

	if err!=nil{
		// Mongo-Error
		if errors.Is(err, mongo.ErrNoDocuments){
			ctx.JSON(http.StatusNotFound, gin.H{
			"error":"⚠️ Note not found for the given ID!",
			"status_code":http.StatusNotFound,
			})
			return
		}

		// Geberal-Error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":"⚠️ Failed to FETCH the note!",
			"status_code":http.StatusInternalServerError,
		})
		return
	}

	// If everything is fine.. ✅
	ctx.JSON(http.StatusOK,note)

}

