package routes

import (
	genrecontroller "github.com/Gswaraw37/goginreactproject/app/controllers/genre_controller"
	moviecontroller "github.com/Gswaraw37/goginreactproject/app/controllers/movie_controller"
	appconfig "github.com/Gswaraw37/goginreactproject/config/app_config"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.Static(appconfig.STATIC_ROUTE, appconfig.STATIC_DIR)

	movieRoute := route.Group("/api/movies")
	movieRoute.GET("/", moviecontroller.Index)
	movieRoute.GET("/:id", moviecontroller.Show)

	genreRoute := route.Group("/api/genre")
	genreRoute.GET("/", genrecontroller.Index)
	genreRoute.GET("/:id/movies", genrecontroller.Show)
}
