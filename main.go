package main

import (
	"flag"
	"fmt"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/handlers"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"os"
	"os/signal"
)

var port = flag.String("port", "8080", "Server port")
var sqlDebug = flag.Bool("sql_debug", false, "Turns on sql debugging")

func main() {
	log.Println("Starting server...")
	flag.Parse()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	router := gin.New()

	srv, err := middleware.NewServer()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.RunMigrationOnDb(srv.Database)
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(srv.Database.DB)
	boil.DebugMode = *sqlDebug

	router.Use(
		middleware.CORSMiddleware,
		srv.DbHandler,
		middleware.RecoveryHandler,
		// middleware.LoggingMiddleware,
	)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})

	v1 := router.Group("/v1")

	// TIME TRIAL
	v1.GET("/time-trials", handlers.GetTimeTrials)
	v1.POST("/time-trials", handlers.CreateTimeTrial)
	v1.PUT("/time-trials", handlers.UpdateTimeTrial)
	v1.GET("/time-trials/:id", handlers.GetTimeTrialById)
	v1.GET("/time-trials/:id/boats", handlers.GetBoatsForTimeTrial)
	v1.GET("/boats/:id", handlers.GetBoatByID)
	v1.POST("/boats", handlers.CreateBoat)
	v1.PUT("/boats", handlers.UpdateBoat)

	lp := fmt.Sprintf(":%s", *port)
	log.WithField("message", fmt.Sprintf("Server Started On Port: %s", *port)).Info()
	router.Run(lp)
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}
