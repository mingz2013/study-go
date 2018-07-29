package room

import "time"

type Room struct {
	Id   int
	Name string

	players []Player // 玩家队列

	// 定义好所有的输入输出接口，就可以定义这个类内部的功能了

	UserMsgIn  <-chan Msg
	UserMsgOut chan<- Msg

	TableMsgIn  <-chan Msg
	TableMsgOut chan<- Msg
}

func NewRoom(id int, userMsgIn <-chan Msg, userMsgOut chan<- Msg, tableMsgIn <-chan Msg, tableMsgOut chan<- Msg) Room {
	return Room{Id: id, Name: "", UserMsgIn: userMsgIn, UserMsgOut: userMsgOut, TableMsgIn: tableMsgIn, TableMsgOut: tableMsgOut}
}

func (r *Room) Init() {

}

func (r *Room) Run() {
	for {
		select {
		case m, ok := <-r.UserMsgIn:
			{
				if !ok {
					continue
				}
				r.DoUserMsg(m)
			}
		case m, ok := <-r.TableMsgIn:
			{
				if !ok {
					continue
				}
				r.DoTableMsg(m)
			}
		case <-time.After(time.Second * 1):
			continue

		}
	}
}

func (r *Room) DoUserMsg(m Msg) {

}

func (r *Room) DoTableMsg(m Msg) {

}

func (r *Room) SendUserRes(Id int, m Msg) {
	m.SetKey("id", Id)
	r.UserMsgOut <- m
}

func (r *Room) SendTableRes(Id int, m Msg) {
	m.SetKey("id", Id)
	r.UserMsgOut <- m
}

func (r *Room) Singin() {
	// 比赛报名

}

func (r *Room) Logout() {
	// 比赛退出
}

func (r *Room) GetInfo() {
	// 获取房间信息
}
