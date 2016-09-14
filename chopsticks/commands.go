package chopsticks

type Command func(p *Player) (string, error)

var Commands = map[string]Command{
	"split": SplitCommand,
	"add":   AddToCommand,
}

func SplitCommand(p *Player) (string, error) {
	return "", nil
}
func AddToCommand(p *Player) (string, error) {
	return "", nil
}

/*
type Command string

const (
	Split Command = "split"
	Add Command = "add"
)

func Parse(s string) (Command, error) {
	if strings.Contains(s, Split) {
		return Split, nil
	}
	if strings.Contains(s, Add) {
		return Add, nil
	}
	return nil, nil
}


func (g *Game) RunCommand(c Command) error {

}
*/
