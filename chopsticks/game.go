package chopsticks

import (
	"fmt"
)

// Chopsticks
// Each player has two hands with up to five fingers on each
// Player moves by "giving" a hand to one of the other player's hands
// This adds that many fingers to that hand, according to modulo arithmetic

const (
	MAX_HAND_COUNT = 5
)

type Move struct {
	giving    *Hand
	receiving *Hand
}

type Game struct {
	Players []*Player
}

func (g Game) String() string {
	out := ""
	for _, p := range g.Players {
		out = fmt.Sprintf("%s %v\n", out, p)
	}
	return out
}

func (g Game) IsOver() bool {
	ret := true
	for _, p := range g.Players {
		ret = ret && p.IsOut()
	}
	return ret
}
