package lib

type Player struct {
	Name   string
	Damage int
}

func NewPlayer(name string, damage int) *Player {
	return &Player{
		Name:   name,
		Damage: damage,
	}
}
