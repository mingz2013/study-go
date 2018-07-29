package table

import (
	"github.com/mingz2013/study.go/test-910-game-server/msg"
	"log"
	"time"
	"math/rand"
)

type Play struct {
	table    Table
	tilePool []int
}

func (p Play) initTilePool() {
	var single []int
	for i := 1; i < 10; i++ {
		single = append(single, i)
	}
	for i := 11; i < 20; i++ {
		single = append(single, i)
	}
	for i := 21; i < 30; i++ {
		single = append(single, i)
	}
	for i := 31; i < 38; i++ {
		single = append(single, i)
	}
	tilePool := make([]int, len(single)*4)
	copy(tilePool, single)
	copy(tilePool[len(single):], single)
	copy(tilePool[len(single)*2:], single)
	copy(tilePool[len(single)*3:], single)

	rand.Seed(time.Now().UnixNano())
	for i := range tilePool {
		j := rand.Intn(i + 1)
		tilePool[i], tilePool[j] = tilePool[j], tilePool[i]

	}

	p.tilePool = tilePool

}

func (p Play) Init() {
	p.initTilePool()
}

func NewPlay(t Table) Play {
	p := Play{table: t}
	p.Init()
	return p
}

func (p Play) OnMsg(m msg.Msg) {

}

func (p Play) Run() {
	log.Println("play run...")
	p.start()
}

func (p Play) start() {
	p.kaiPai()
	p.nextOp(0)
}

func (p Play) kaiPai() {

	for i := 0; i < 4; i++ {
		kaiPai := p.tilePool[:13]
		p.tilePool = p.tilePool[13:]

		p.table.Players[i].Cards.DoKaiPai(kaiPai)

	}

	p.SendKaiPaiRes()

}

func (p Play) SendKaiPaiRes() {
	for i := 0; i < 4; i++ {

		//player := p.table.Players[i]

	}
}


func (p Play) nextOp(seatId int) {

}
