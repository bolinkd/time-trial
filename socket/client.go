package socket

import (
	"github.com/bolinkd/time-trial/socket/model"
	"github.com/pkg/errors"
)

type ClientInterface interface {
	SendUpdateEventByUser(userId int64, updateEvent model.UpdateEvent) error
	SendUpdateEventByChannel(channelId string, updateEvent model.UpdateEvent) error
	SendUpdateEventByRoom(roomName string, updateEvent model.UpdateEvent)
}

type Client struct{}

func New() Client {
	return Client{}
}

func (client *Client) SendUpdateEventByUser(userId int64, updateEvent model.UpdateEvent) error {
	return errors.New("NOT IMPLEMENTED")
}

func (client *Client) SendUpdateEventByChannel(channelId string, updateEvent model.UpdateEvent) error {
	channel, err := socketServer.GetChannel(channelId)
	if err != nil {
		return err
	}

	return channel.Emit(updateEvent.Type, updateEvent.Payload)
}

func (client *Client) SendUpdateEventByRoom(roomName string, updateEvent model.UpdateEvent) {
	socketServer.BroadcastTo(roomName, updateEvent.Type, updateEvent.Payload)
}
