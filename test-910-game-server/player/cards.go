package player

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
