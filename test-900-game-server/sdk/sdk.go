package sdk

import (
	"github.com/mingz2013/study.go/test-900-game-server/conf"
	"github.com/mingz2013/study.go/test-900-game-server/database"
)

type SDK struct{}

type LoginReq struct {
	DeviceId string
}

type LoginRes struct {
	ServerAddr conf.ServerAddr
	User       database.User
}
