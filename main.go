package main

import (
	"log"
	"net/http"
	"seatmap-backend/api/handler"
	"seatmap-backend/api/middlewares"
	"seatmap-backend/api/routes"
	"seatmap-backend/infrastructure/repository"
	"seatmap-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

type EnvConfig struct {
	MigrationURL string `mapstructure:"MIGRATION_URL"`
	DBSource     string `mapstructure:"DB_SOURCE"`
	Port         string `mapstructure:"PORT"`
}

func main() {
	config, err := loadEnv(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	repository.ConnectDB(config.DBSource)
	runDBMigration(config.MigrationURL, config.DBSource)
	r := SetupRouter()
	_ = r.Run(":8080")
}

func loadEnv(path string) (config EnvConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	// r.Use(middlewares.JSONAppErrorReporter())
	db := repository.GetDB()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	roleRepo := repository.NewRoleRepo(db)
	roleService := services.NewRoleService(roleRepo)
	rolehandler := handler.NewRoleHandler(roleService)
	roleRoutes := routes.NewRoleRoutes(rolehandler)
	roleRoutes.Setup(r)

	userRepo := repository.NewUserRepo(db)
	userService := services.NewUserService(userRepo, roleService)
	userhandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userhandler)
	userRoutes.Setup(r)

	return r
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cant create new migrate instance: ", err)
	}
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("fail to run migrate up: ", err)
	}
	log.Println("db migrate successfully")
}