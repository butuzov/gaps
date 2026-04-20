package gaps

type Mode uint64

const (
	ModeDefault Mode = 1 << iota
	ModeFireAndForgot
	ModeWatch
	ModeAnalize
	ModeCompare
)

func (m Mode) Has(flag Mode) bool {
	return m&flag != 0
}

func (m Mode) HasAll(flags Mode) bool {
	return m&flags == flags
}
