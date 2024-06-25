package moviecontroller

import (
	"github.com/Gswaraw37/goginreactproject/app/models"
	"github.com/Gswaraw37/goginreactproject/database"
	"github.com/Gswaraw37/goginreactproject/responses"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	movies := new([]models.Movie)
	err := database.DB.Table("movies").Preload("MovieGenre.Genre").Find(&movies).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    movies,
	})
}

func Show(ctx *gin.Context) {
	id := ctx.Param("id")
	movie := new(responses.MovieResponse)

	err := database.DB.Table("movies").Preload("MovieGenre.Genre").Where("id = ?", id).Find(&movie, id).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if movie.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    movie,
	})
}
