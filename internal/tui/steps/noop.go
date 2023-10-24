package steps

type Noop struct {
}

func NewNoopStep() Noop {
	return Noop{}
}

func (t Noop) Filled() bool {
	return false
}
