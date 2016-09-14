package chopsticks

type invalidSplitError struct {
	S string
}

func (e *invalidSplitError) Error() string {
	return e.S
}
