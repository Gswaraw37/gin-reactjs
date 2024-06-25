package responses

import (
	"time"

	"github.com/Gswaraw37/goginreactproject/app/models"
)

type MovieResponse struct {
	ID          *int                `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Year        int                 `json:"year"`
	ReleaseDate time.Time           `json:"release_date"`
	Runtime     int                 `json:"runtime"`
	Rating      int                 `json:"rating"`
	MPAARating  string              `json:"mpaa_rating"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	MovieGenre  []models.MovieGenre `json:"genres" gorm:"foreignKey:MovieID;references:ID"`
}

func (MovieResponse) TableName() string {
	return "movies"
}
