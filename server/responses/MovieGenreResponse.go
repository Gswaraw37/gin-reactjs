package responses

import "github.com/Gswaraw37/goginreactproject/app/models"

type MovieGenreResponse struct {
	ID    *int           `json:"id"`
	Genre []models.Genre `json:"genre" gorm:"foreignKey:GenreID;references:ID"`
}

func (MovieGenreResponse) TableName() string {
	return "movies_genres"
}
