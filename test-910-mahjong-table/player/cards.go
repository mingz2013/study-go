package player

import "github.com/mingz2013/study.go/test-910-mahjong-table/actions"

type HandPile struct {
	tiles []int
}

type DropPile struct {
	tiles []int
}

type Pile struct {
}

type ChiPile struct {
}

type PengPile struct {
}

type GangPile struct {
}

type Cards struct {
	handPile HandPile
	dropPile DropPile
	cpgPiles []Pile
	nowTile  int
}

func (c Cards) Init() {
	c.handPile = HandPile{}
	c.dropPile = DropPile{}
}

func NewCards() Cards {
	return Cards{}
}

func (c Cards) DoKaiPaiAction(action actions.KaiPaiAction) {
	copy(c.handPile.tiles, action.Tiles)
}

