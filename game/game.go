package main

import (
	"fmt"
	"github.com/billy-hardy/ic-weiner/gameErrors"
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

type Hand struct {
	count int
}

type Player struct {
	Name        string
	right, left *Hand
}

type Game struct {
	players []Player
}

func (g Game) String() string {
	out := ""
	for _, p := range g.players {
		out = fmt.Sprintf("%s %v\n", out, p)
	}
	return out
}

func (g Game) IsOver() bool {
	ret := true
	for _, p := range g.players {
		ret = ret && p.IsOut()
	}
	return ret
}

func (p Player) String() string {
	return fmt.Sprintf("%s's state:\n\tleft hand: %v\n\tright hand: %v", p.Name, p.left, p.right)
}

func (h Hand) String() string {
	return fmt.Sprintf("%d", h.count)
}

func (p Player) IsOut() bool {
	return p.right.IsOut() && p.left.IsOut()
}

func (h *Hand) AddTo(g Game, h2 *Hand) {
	h2.count = (h2.count + h.count) % MAX_HAND_COUNT
}

func (p *Player) Split() error {
	if !p.right.IsOut() && !p.left.IsOut() {
		return &gameErrors.InvalidSplitError{"Can not split without an empty hand"}
	}
	if p.right.IsOut() {
		return split(p.left, p.right)
	}
	return split(p.right, p.left)
}

func (h Hand) IsOut() bool {
	return h.count == 0
}

func split(toSplit *Hand, empty *Hand) error {
	if toSplit.count%2 != 0 {
		return &gameErrors.InvalidSplitError{"Hand being split must have an even number"}
	}
	toSplit.count /= 2
	empty.count = toSplit.count
	return nil
}

func main() {
	players := []Player{Player{
		"player one",
		&Hand{0},
		&Hand{4},
	}, Player{
		"player two",
		&Hand{3},
		&Hand{1},
	}}
	game := Game{
		players,
	}
	repl()
	printToScreen(game)
	printToScreen(game.players[0].Split())
	printToScreen(game)
	printToScreen(game.players[0].Split())
	printToScreen(game)
}

type Command func(p *Player) error

func repl() {
	commands := map[string]Command{
		"split": SplitCommand,
	}
	commands["add"] = AddToCommand
}

func printToScreen(obj interface{}) {
	if obj != nil {
		fmt.Println(obj)
	}
}

func SplitCommand(p *Player) error {
	return nil
}
func AddToCommand(p *Player) error {
	return nil
}
