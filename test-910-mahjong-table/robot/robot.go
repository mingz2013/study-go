package robot

import (
	"github.com/mingz2013/study.go/test-910-mahjong-table/msg"
	"time"
	"os"
	"log"
)

type Robot struct {
	Id     int
	Name   string
	MsgIn  <-chan msg.Msg
	MsgOut chan<- msg.Msg
}

func NewRobot(id int, name string, msgIn <-chan msg.Msg, msgOut chan<- msg.Msg) Robot {
	return Robot{Id: id, Name: name, MsgIn: msgIn, MsgOut: msgOut}
}

func (r Robot) doSit() {

	//m := msg.Msg{"cmd":"sit", "params":map[string]interface{}{"id": r.Id, "name": r.Name}}
	m := msg.NewMsg()
	m.SetCmd("sit")
	m.SetParams(map[string]interface{}{"id": r.Id, "name": r.Name})
	r.MsgOut <- m

}

func (r Robot) Run() {

	r.doSit()

	for {

		select {
		case m, ok := <-r.MsgIn:
			{
				if !ok {
					continue
				}

				r.onMsg(m)
			}
		case <-time.After(1 * time.Second):
			continue
		}

	}
}

func (r Robot) onMsg(m msg.Msg) {
	switch m.GetCmd() {
	case "sit":
		{
			params := m.GetResults()
			retCode := params["retcode"].(int)
			msgRet := params["msg"].(string)
			if retCode != 0 {
				log.Println(msgRet)
				os.Exit(retCode)
			}
			log.Println(msgRet)

		}

	default:
		log.Println("unknown msg", m)

	}
}
