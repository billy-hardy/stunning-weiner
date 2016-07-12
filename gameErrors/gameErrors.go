package gameErrors

type InvalidSplitError struct {
	S string
}

func (e *InvalidSplitError) Error() string {
	return e.S
}
