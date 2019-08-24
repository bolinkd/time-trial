package main

import (
	"flag"
	"fmt"
	"github.com/bolinkd/time-trial/db"
	"github.com/bolinkd/time-trial/handlers"
	"github.com/bolinkd/time-trial/middleware"
	"github.com/bolinkd/time-trial/socket"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
	"os"
	"os/signal"
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
		// middleware.LoggingMiddleware,
	)

	socket.InitSocketServer(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})

	v1 := router.Group("/v1")

	// ORGANIZATION
	v1.GET("/organizations", handlers.GetOrganizations)
	v1.GET("/organizations/:id", handlers.GetOrganizationByID)
	v1.POST("/organizations", handlers.CreateOrganization)
	v1.PUT("/organizations", handlers.UpdateOrganization)
	v1.GET("/organizations/:id/clubs", handlers.GetClubsByOrganization)

	// CLUB
	v1.GET("/clubs/:id", handlers.GetClubByID)
	v1.POST("/clubs", handlers.CreateClub)
	v1.PUT("/clubs", handlers.UpdateClub)
	v1.GET("/clubs/:id/shells", handlers.GetShellsByClub)
	v1.GET("/clubs/:id/groups", handlers.GetGroupsByClub)

	// SHELL
	v1.GET("/shells/:id", handlers.GetShellByID)
	v1.POST("/shells", handlers.CreateShell)
	v1.PUT("/shells", handlers.UpdateShell)
	v1.GET("/shells/:id/rentals", handlers.GetRentalsByShell)

	// GROUP (ie adult, junior, novice)
	v1.GET("/groups/:id", handlers.GetGroupByID)
	v1.POST("/groups", handlers.CreateGroup)
	v1.PUT("/groups", handlers.UpdateGroup)
	v1.GET("/groups/:id/rowers", handlers.GetRowersByGroup)

	// ROWER
	v1.GET("/rowers/:id", handlers.GetRowerByID)
	v1.POST("/rowers", handlers.CreateRower)
	v1.PUT("/rowers", handlers.UpdateRower)

	// RENTAL
	v1.GET("/rentals", handlers.GetRentals)
	v1.GET("/rentals/:id", handlers.GetRentalByID)
	v1.POST("/rentals", handlers.CreateRental)
	v1.PUT("/rentals", handlers.UpdateRental)
	v1.DELETE("/rentals/:id", handlers.DeleteRental)
	v1.GET("/rentals/:id/rowers", handlers.GetRentalRowersByRental)

	// RENTAL ROWERS
	v1.GET("/rental-rowers/:id", handlers.GetRentalRowerByID)
	v1.POST("/rental-rowers", handlers.CreateRentalRower)
	v1.PUT("/rental-rowers", handlers.UpdateRentalRower)
	v1.DELETE("/rental-rowers/:id", handlers.DeleteRentalRower)

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
