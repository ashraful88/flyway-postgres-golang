package main

import (
	"log"
	"net/http"
	"os"

	"flyway-postgres-golang/flyway"

	"github.com/gin-gonic/gin"
)

func main() {
	env, _ := os.LookupEnv("ENV")
	ApiPort, _ := os.LookupEnv("API_PORT")
	dbHost, _ := os.LookupEnv("POSTGRES_HOST")
	dbPort, _ := os.LookupEnv("POSTGRES_PORT")
	dbName, _ := os.LookupEnv("POSTGRES_DB")
	dbUser, _ := os.LookupEnv("POSTGRES_USER")
	dbPass, _ := os.LookupEnv("POSTGRES_PASSWORD")

	if env != "local" {
		cs := flyway.CreateConnectionString(dbHost, dbPort, dbName, dbUser, dbPass, "filesystem:/flyway/sql", 60)

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

	r.GET("/health", HealthCheck)

	err := r.Run(":" + ApiPort)
	if err != nil {
		panic(err)
	}
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
