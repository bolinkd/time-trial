package domain

type M struct {
	Message string `json:"message"`
}

func Message(message string) M {
	return M{
		Message: message,
	}
}
