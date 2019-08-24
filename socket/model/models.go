package model

type AuthEvent struct {
	Token string `json:"token"`
}

type Event struct {
	UpdateMethod string      `json:"update_method"`
	DataType     string      `json:"data_type"`
	Payload      interface{} `json:"payload"`
}

type UpdateEvent struct {
	Type    string `json:"type"`
	Payload Event  `json:"payload"`
}

type ChannelUpdateMessage struct {
	ChannelID string `json:"channel_id"`
	UpdateEvent
}

type UserUpdateMessage struct {
	UserID      int64 `json:"user_id"`
	UpdateEvent `json:"update_event"`
}

type JoinRaceEvent struct {
	UpdateMethod string          `json:"update_method"`
	DataType     string          `json:"data_type"`
	Payload      JoinRaceMessage `json:"payload"`
}

type JoinRaceMessage struct {
	Timestamp int64  `json:"timestamp"`
	RaceId    int64  `json:"race_id"`
	Timers    int    `json:"timers"`
	DeviceId  string `json:"device_id"`
}
