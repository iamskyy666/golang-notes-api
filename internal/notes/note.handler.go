package notes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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