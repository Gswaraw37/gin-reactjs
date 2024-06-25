package responses

import "time"

type GenreResponse struct {
	ID        *int      `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (GenreResponse) TableName() string {
	return "genres"
}
