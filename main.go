package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/handlers"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/bolinkd/time-trial/socket"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

var flagPort = flag.String("port", "8080", "Server port")
var sqlDebug = flag.Bool("sql_debug", false, "Turns on sql debugging")

func main() {
	log.Println("Starting server...")
	flag.Parse()

	port := os.Getenv("PORT")
	if port == "" {
		port = *flagPort
	}

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
		srv.DarkSkyHandler,
		srv.SocketServiceHandler,
		srv.ServicesHandler,
		middleware.RecoveryHandler,
		middleware.LoggingMiddleware,
	)

	socket.InitSocketServer(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})

	v1 := router.Group("/v1")

	// ORGANIZATION
	v1.GET("/organizations", handlers.GetOrganizations)
	auth := v1.Group("/").Use(middleware.AuthMiddleware)
	auth.GET("organization", handlers.GetCurrentOrganization)
	auth.GET("/organizations/:id", handlers.GetOrganizationByID)
	auth.POST("/organizations", handlers.CreateOrganization)
	auth.PUT("/organizations", handlers.UpdateOrganization)

	// SHELL
	auth.GET("/shells", handlers.GetShellsByCurrentOrganization)
	auth.GET("/shells/:id", handlers.GetShellByID)
	auth.POST("/shells", handlers.CreateShell)
	auth.PUT("/shells", handlers.UpdateShell)
	auth.GET("/shells/:id/rentals", handlers.GetRentalsByShell)

	// GROUP (ie adult, junior, novice)
	auth.GET("/groups", handlers.GetGroupsByCurrentOrganization)
	auth.GET("/groups/:id", handlers.GetGroupByID)
	auth.POST("/groups", handlers.CreateGroup)
	auth.PUT("/groups", handlers.UpdateGroup)
	auth.GET("/groups/:id/rowers", handlers.GetRowersByGroup)
	auth.GET("/groups/:id/shells", handlers.GetShellsByGroup)

	// ROWER
	auth.GET("/rowers", handlers.GetRowersByCurrentOrganization)
	auth.GET("/rowers/:id", handlers.GetRowerByID)
	auth.POST("/rowers", handlers.CreateRower)
	auth.PUT("/rowers", handlers.UpdateRower)

	// RENTAL
	auth.GET("/rentals", handlers.GetRentals)
	auth.GET("/rentals/:id", handlers.GetRentalByID)
	auth.POST("/rentals", handlers.CreateRental)
	auth.PUT("/rentals", handlers.UpdateRental)
	auth.DELETE("/rentals/:id", handlers.DeleteRental)
	auth.GET("/rentals/:id/rowers", handlers.GetRentalRowersByRental)

	// RENTAL ROWERS
	auth.GET("/rental-rowers/:id", handlers.GetRentalRowerByID)
	auth.POST("/rental-rowers", handlers.CreateRentalRower)
	auth.DELETE("/rental-rowers/:id", handlers.DeleteRentalRower)

	// TIME TRIAL
	v1.GET("/time-trials", handlers.GetTimeTrials)
	v1.POST("/time-trials", handlers.CreateTimeTrial)
	v1.PUT("/time-trials", handlers.UpdateTimeTrial)
	v1.GET("/time-trials/:id", handlers.GetTimeTrialById)
	v1.GET("/time-trials/:id/boats", handlers.GetBoatsForTimeTrial)

	// BOAT
	v1.GET("/boats/:id", handlers.GetBoatByID)
	v1.POST("/boats", handlers.CreateBoat)
	v1.PUT("/boats", handlers.UpdateBoat)
	v1.GET("/weather", handlers.AddWeatherToTimeTrial)

	v1.POST("/authenticate", handlers.Authenticate)

	lp := fmt.Sprintf(":%s", port)
	log.WithField("message", fmt.Sprintf("Server Started On Port: %s", port)).Info()
	if err := http.ListenAndServe(lp, router); err != nil {
		log.Panic(err.Error())
	}
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}
