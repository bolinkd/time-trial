package main

import (
	"flag"
	"fmt"
	"github.com/braintree/manners"
	"github.com/businessinstincts/traxone/handlers"
	"github.com/businessinstincts/traxone/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
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

	setup.ConfigureLogging(log.StandardLogger())

	err = db.RunMigrationOnDb(srv.Database)
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(srv.Database.DB)
	boil.DebugMode = *sqlDebug

	router.Use(
		srv.AppNexusServiceHandler,
		middleware.CORSMiddleware,
		srv.OnspotServiceHandler,
		srv.SocketServiceHandler,
		srv.DbHandler,
		middleware.RecoveryHandler,
		middleware.LoggingMiddleware,
		srv.PayTraceHandler,
	)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})

	v1 := router.Group("/v1")
	v1.Use(
		middleware.JWTSPAMiddleware,
	)
	machine := router.Group("/machine")
	machine.Use(
		middleware.JWTMachineMiddleware,
	)
	onspot := router.Group("/public/v1/onspot")

	//COUNT
	onspot.POST("/count", handlers.GetDeviceCountByGeoframe)
	onspot.POST("/count/callback", handlers.GetDeviceCountByGeoframeCallback)

	//USER
	v1.GET("/users/:id", handlers.GetUserByID)
	machine.POST("/users", handlers.CreateUser)
	v1.POST("/users", handlers.CreateUser)
	v1.GET("/users", handlers.GetUser)
	v1.PUT("/users", handlers.UpdateUser)

	//ORG
	v1.GET("/organizations/:id", handlers.GetOrgByID)
	v1.POST("/organizations", handlers.CreateOrg)
	v1.GET("/organizations", handlers.GetOrgByUserID)
	v1.PUT("/organizations", handlers.UpdateOrg)

	//CAMPAIGN
	v1.GET("/users/:id/campaigns", handlers.GetCampaignsByUserID)
	v1.GET("/campaigns", handlers.GetCampaigns)
	v1.GET("/campaigns/:id", handlers.GetCampaignByID)
	v1.POST("/campaigns", handlers.CreateCampaign)
	//v1.PUT("/campaigns", handlers.UpdateCampaign)

	admin := v1.Group("/admin")
	admin.PUT("/campaigns", handlers.UpdateCampaign)
	admin.GET("/campaigns", handlers.GetAllCampaigns)

	lp := fmt.Sprintf(":%s", *port)
	log.WithField("message", fmt.Sprintf("Server Started On Port: %s", *port)).Info()
	if err := http.ListenAndServe(lp, router); err != nil {
		log.Panic(err.Error())
	}
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}
