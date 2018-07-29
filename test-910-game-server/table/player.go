package table

type Player struct {
	Id   int
	name string

	SeatId int
}

func NewPlayer(seatId int) Player {
	return Player{SeatId: seatId}
}
