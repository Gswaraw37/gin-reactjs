package requests

import (
	"time"

	"github.com/Gswaraw37/goginreactproject/app/models"
)

type MovieRequest struct {
	Title       string    `json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	Year        int       `json:"year" form:"year" binding:"required"`
	ReleaseDate time.Time `json:"release_date" form:"release_date" binding:"required"`
	Runtime     int       `json:"runtime" form:"runtime" binding:"required"`
	Rating      int       `json:"rating" form:"rating" binding:"required"`
	MPAARating  string    `json:"mpaa_rating" form:"mpaa_rating" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" binding:"required"`
	MovieGenre  []models.MovieGenre
}
