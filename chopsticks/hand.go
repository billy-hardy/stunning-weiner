package chopsticks

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
)

type Hand struct {
	*js.Object
	Count int
}

func (h Hand) String() string {
	return fmt.Sprintf("%d", h.Count)
}
func (h *Hand) AddTo(g Game, h2 *Hand) {
	h2.Count = (h2.Count + h.Count) % MAX_HAND_COUNT
}

func (h Hand) IsOut() bool {
	return h.Count == 0
}
