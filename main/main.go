package main

import (
	"bufio"
	"fmt"
	"github.com/billy-hardy/ic-weiner/chopsticks"
	"os"
)

func main() {
	chopsticks.SetGlobals()
	players := []*chopsticks.Player{
		&chopsticks.Player{
			Name:  "Player one",
			Right: &chopsticks.Hand{Count: 0},
			Left:  &chopsticks.Hand{Count: 4},
		},
		&chopsticks.Player{
			Name:  "Player two",
			Right: &chopsticks.Hand{Count: 3},
			Left:  &chopsticks.Hand{Count: 1},
		},
	}
	game := chopsticks.Game{
		Players: players,
	}
	sendResponse(game)
	sendResponse(game.Players[0].Split())
	sendResponse(game)
	sendResponse(game.Players[0].Split())
	sendResponse(game)
	repl(game)
}

func repl(g chopsticks.Game) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := getInput(reader)
		if err != nil {
			sendResponse(err)
			break
		}
		//		chopsticks.Parse(input)
		sendResponse(input)
		sendResponse(chopsticks.Commands["split"](g.Players[0]))
	}
}

func getInput(reader *bufio.Reader) (string, error) {
	return reader.ReadString('\n')
}

func sendResponse(obj ...interface{}) {
	for _, val := range obj {
		if val != nil {
			fmt.Println(val)
		}
	}
}
