package bootstrap

import (
	"log"

	"github.com/Gswaraw37/goginreactproject/config"
	appconfig "github.com/Gswaraw37/goginreactproject/config/app_config"
	corsconfig "github.com/Gswaraw37/goginreactproject/config/cors_config"
	logconfig "github.com/Gswaraw37/goginreactproject/config/log_config"
	"github.com/Gswaraw37/goginreactproject/database"
	"github.com/Gswaraw37/goginreactproject/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.InitConfig()
	database.ConnectDatabase()
	logconfig.DefaultLogger()

	app := gin.Default()
	app.Use(corsconfig.CorsConfig())

	routes.InitRoute(app)
	app.Run(":" + appconfig.PORT)
}
