package socket

import (
	"encoding/json"
	"fmt"
	"github.com/bolinkd/time-trial/models"
	"github.com/bolinkd/time-trial/socket/model"
	"github.com/gin-gonic/gin"
	"github.com/mtfelian/golang-socketio"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"strconv"
	"time"
)

const (
	EventConnected = "connected"
	EventError     = "error"
	EventJoinRace  = "join-race"

	RoomRaceData       = "race-data"
	RoomJoinRacePrefix = "race-"
)

var socketServer *gosocketio.Server

func init() {
	socketServer = gosocketio.NewServer()
}

func onConnectionHandler(c *gosocketio.Channel) error {
	cid := c.Id()
	log.Printf("Connected %s\n", cid)
	err := c.Join(RoomRaceData)
	if err != nil {
		return c.Emit(EventError, err.Error())
	}
	return c.Emit(EventConnected, "connected")
}

func onDisconnectionHandler(c *gosocketio.Channel) error {
	cid := c.Id()
	log.Printf("Disconnected %s\n", cid)
	return c.Leave(RoomRaceData)
}

func onJoinRaceHandler(c *gosocketio.Channel, param interface{}) error {
	j, err := json.Marshal(param)
	if err != nil {
		log.Println("error:", err)
		return err
	}
	log.Printf("payload JSON is %s\n", j)

	var obj model.JoinRaceEvent
	if err := json.Unmarshal(j, &obj); err != nil {
		log.Println("error:", err)
		return err
	}
	roomName := RoomJoinRacePrefix + strconv.FormatInt(obj.Payload.RaceId, 10)
	err = c.Join(roomName)
	if err != nil {
		log.WithError(err).Error(err.Error())
		return err
	}

	timeTrial, err := models.TimeTrials(qm.Where("id = ?", obj.Payload.RaceId)).One(boil.GetDB())
	if err != nil {
		return err
	}
	fmt.Println(roomName, timeTrial.Timers, timeTrial.ID, timeTrial.StartTime, obj.Payload.Timestamp, time.Now().Unix(), c.Id(), len(c.List(roomName)))
	if len(c.List(roomName)) == obj.Payload.Timers {
		c.BroadcastTo(roomName, roomName, "time-request")
		// if err := socketServer.On(roomName, startRaceHandler); err != nil {
		// log.Fatal(err)
		// }
	}
	return nil
}

func InitSocketServer(router *gin.Engine) {
	if err := socketServer.On(gosocketio.OnConnection, onConnectionHandler); err != nil {
		log.Fatal(err)
	}
	if err := socketServer.On(gosocketio.OnDisconnection, onDisconnectionHandler); err != nil {
		log.Fatal(err)
	}
	if err := socketServer.On(EventJoinRace, onJoinRaceHandler); err != nil {
		log.Fatal(err)
	}

	log.Print("Socket server init")

	router.Any("/socket.io/", gin.WrapH(socketServer))
}
