package table

import "github.com/mingz2013/study.go/test-910-game-server/msg"

type Play struct {
}

func NewPlay() Play {
	return Play{}
}

func (p Play) OnMsg(m msg.Msg) {

}

func (p Play) Run() {
	p.start()
}

func (p Play) start() {
	p.kaiPai()
	p.nextOp(0)
}

func (p Play) kaiPai() {

}

func (p Play) nextOp(seatId int) {

}
