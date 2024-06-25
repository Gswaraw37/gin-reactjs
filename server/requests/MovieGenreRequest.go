package requests

import (
	"time"

	"github.com/Gswaraw37/goginreactproject/app/models"
)

type MovieGenreRequest struct {
	MovieID   int          `json:"movie_id" form:"movie_id" binding:"required"`
	GenreID   int          `json:"genre_id" form:"genre_id" binding:"required"`
	Genre     models.Genre `json:"genre" form:"genre" binding:"required"`
	CreatedAt time.Time    `json:"created_at" form:"created_at" binding:"required"`
	UpdatedAt time.Time    `json:"updated_at" form:"updated_at" binding:"required"`
}
