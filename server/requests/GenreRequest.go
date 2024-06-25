package requests

import "time"

type GenreRequest struct {
	GenreName string    `json:"genre_name" form:"genre_name" binding:"required"`
	CreatedAt time.Time `json:"created_at" form:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" binding:"required"`
}
