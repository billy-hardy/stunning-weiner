package chopsticks

import "github.com/gopherjs/gopherjs/js"

func NewGame(players []*Player) *js.Object {
	return js.MakeWrapper(&Game{Players: players})
}

func NewPlayer(name string, right, left *Hand) *js.Object {
	return js.MakeWrapper(&Player{Name: name, Right: right, Left: left})
}

func NewHand(count int) *js.Object {
	return js.MakeWrapper(&Hand{Count: count})
}

func SetGlobals() {
	js.Global.Set("chopsticks", map[string]interface{}{
		"Hand":   NewHand,
		"Player": NewPlayer,
		"Game":   NewGame,
	})
}
