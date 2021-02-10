package sdk

import (
	"study-go/test-900-game-server/conf"
	"study-go/test-900-game-server/database"
)

type SDK struct{}

type LoginReq struct {
	DeviceId string
}

type LoginRes struct {
	ServerAddr conf.ServerAddr
	User       database.User
}
