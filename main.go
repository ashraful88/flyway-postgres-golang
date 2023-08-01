package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ashraful88/flyway-postgres-golang/config"
	"github.com/ashraful88/flyway-postgres-golang/flyway"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	dbHost, _ := os.LookupEnv("POSTGRES_HOST")
	dbPort, _ := os.LookupEnv("POSTGRES_PORT")
	dbName, _ := os.LookupEnv("POSTGRES_DB")
	dbUser, _ := os.LookupEnv("POSTGRES_USER")
	dbPass, _ := os.LookupEnv("POSTGRES_PASSWORD")

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading primary .env file")
	}
	// get env from .env file
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	// load configs non based on env
	config.Conf = config.GetConfigByEnv(env)

	if env != "local" {
		cs := flyway.CreateConnectionString(dbHost, dbPort, dbName, dbUser, dbPass, config.Conf.FlywayDBSchemas, config.Conf.FlywayLocations, config.Conf.FlywayConnRetry)

		// wipe all database schema and data
		/* res, err := flyway.Clean(cs)
		if err != nil {
			log.Println("Flyway DB Clean Failed", err)
		}
		log.Println("flyway Clean", res) */

		// run migration
		res, err := flyway.Migrate(cs)
		if err != nil {
			log.Println("Flyway DB Migration Failed", err)
		}

		log.Println("flyway Migrate", res)
	}

	r := gin.New()
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.Use(func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Methods", "*")
		context.Writer.Header().Add("Access-Control-Allow-Headers", "*")
		context.Writer.Header().Set("Content-Type", "application/json; charset=latin-1")
		context.Next()
	})
	_ = r.SetTrustedProxies(nil)

	r.GET("/health", HealthCheck)

	err = r.Run(":" + config.Conf.APIPort)
	if err != nil {
		panic(err)
	}
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
