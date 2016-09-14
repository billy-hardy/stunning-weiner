package chopsticks

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
)

type Player struct {
	*js.Object
	Name        string
	Right, Left *Hand
}

func (p Player) String() string {
	return fmt.Sprintf("%s's state:\n\tleft hand: %v\n\tright hand: %v", p.Name, p.Left, p.Right)
}

func (p Player) IsOut() bool {
	return p.Right.IsOut() && p.Left.IsOut()
}

func (p *Player) Split() error {
	if !p.Right.IsOut() && !p.Left.IsOut() {
		return &invalidSplitError{"Can not split without an empty hand"}
	}
	if p.Right.IsOut() {
		return split(p.Left, p.Right)
	}
	return split(p.Right, p.Left)
}

func split(toSplit *Hand, empty *Hand) error {
	if toSplit.Count%2 != 0 {
		return &invalidSplitError{"Hand being split must have an even number"}
	}
	toSplit.Count /= 2
	empty.Count = toSplit.Count
	return nil
}
