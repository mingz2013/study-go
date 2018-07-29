package room

type RoomSession struct {
	RoomId      int
	Room        Room
	UserMsgIn   chan<- Msg
	UserMsgOut  <-chan Msg
	TableMsgIn  chan<- Msg
	TableMsgOut <-chan Msg
}

type RoomManager struct {
	RoomSessionMap map[int]RoomSession

	MsgIn  chan Msg
	MsgOut chan Msg
}

func (rm *RoomManager) CreateNewRoom(roomId int) RoomSession {
	roomSession := RoomSession{}
	roomSession.RoomId = roomId

	userMsgIn := make(chan Msg)
	userMsgOut := make(chan Msg)
	tableMsgIn := make(chan Msg)
	tableMsgOut := make(chan Msg)

	roomSession.Room = NewRoom(roomId, userMsgIn, userMsgOut, tableMsgIn, tableMsgOut)
	roomSession.UserMsgIn = userMsgIn
	roomSession.UserMsgOut = userMsgOut
	roomSession.TableMsgIn = tableMsgIn
	roomSession.UserMsgOut = tableMsgOut

	rm.RoomSessionMap[roomId] = roomSession

	return roomSession

}
