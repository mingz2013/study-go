package table

import (
	"github.com/mingz2013/study.go/test-910-game-server/msg"
	"log"
	"time"
)

type Table struct {
	Id string // 桌子Id，由manager自动生成

	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg

	Play    Play
	Players [4]Player
}

func NewTable(id string, msgIn <-chan msg.Msg, msgOut chan<- msg.Msg) Table {
	t := Table{Id: id, MsgIn: msgIn, MsgOut: msgOut}
	t.Init()
	return t
}

func (t Table) Init() {
	for i := 0; i < 4; i++ {
		t.Players[i] = NewPlayer(i)
	}
	t.Play = NewPlay()

}

func (t Table) Run() {

	for {
		select {
		case m, ok := <-t.MsgIn:
			{
				if !ok {
					continue
				}

				t.onMsg(m)
			}
		case <-time.After(1 * time.Second):
			continue

		}

	}
}

func (t Table) onMsg(m msg.Msg) {
	switch m.GetCmd() {
	case "sit":
		t.onCmdSit(m)
	case "play":
		t.Play.OnMsg(m)
	default:
		log.Println("unknown cmd", m)
	}
}

func (t Table) getBestSeatId() (int, bool) {
	for i := 0; i < 4; i++ {
		if t.Players[i].Id == 0 {
			return i, true
		}
	}
	return -1, false
}

func (t Table) onCmdSit(m msg.Msg) {
	params := m.GetParams()
	id := params["id"].(int)
	name := params["name"].(string)

	seatId, ok := t.getBestSeatId()

	if !ok {
		log.Println("not found empty seat", id, name)
		t.MsgOut <- msg.Msg{"id": id, "cmd": "sit", "results": map[string]interface{}{"retcode": "-1", "msg": "not found empty seat"}}
		return
	}

	p := t.Players[seatId]
	p.Id = id
	p.name = name

	log.Println("sit ok", id, name, seatId)

	t.MsgOut <- msg.Msg{"id": id, "cmd": "sit", "results": map[string]interface{}{"retcode": "0", "msg": "sit ok"}}
}

func (t Table) checkFull() {
	_, ok := t.getBestSeatId()
	if !ok {
		t.onTableStart()
	}
}

func (t Table) onTableStart() {
	t.Play.Run()
}
