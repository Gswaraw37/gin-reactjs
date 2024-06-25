package genrecontroller

import (
	"github.com/Gswaraw37/goginreactproject/app/models"
	"github.com/Gswaraw37/goginreactproject/database"
	"github.com/Gswaraw37/goginreactproject/responses"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	genres := new([]models.Genre)
	err := database.DB.Table("genres").Find(&genres).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    genres,
	})
}

func Show(ctx *gin.Context) {
	genreId := ctx.Param("id")
	genre := new(responses.GenreResponse)
	movies := new([]models.Movie)

	errGenre := database.DB.Table("genres").Where("id = ?", genreId).Find(&genre, genreId).Error
	if errGenre != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if genre.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	errMovie := database.DB.Table("movies").Joins("JOIN movies_genres ON movies.id = movies_genres.movie_id").Preload("MovieGenre.Genre").Where("movies_genres.genre_id = ?", genreId).Find(&movies).Error
	if errMovie != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"movies":  movies,
		"genre":   genre,
	})
}
