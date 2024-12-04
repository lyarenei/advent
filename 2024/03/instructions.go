package _3

type instruction interface {
	Execute() interface{}
}

type mul struct {
	a int
	b int
}

func (m mul) Execute() interface{} {
	return uint64(m.a) * uint64(m.b)
}

type do struct {
	// Do instruction
}

func (d do) Execute() interface{} {
	return nil
}

type dont struct {
	// Don't instruction
}

func (d dont) Execute() interface{} {
	return nil
}
