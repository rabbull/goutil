package goutil

type Adder interface {
	Incr()
	Clone() Adder
}

type AdderMapper struct {
	adder Adder
	dict  map[interface{}]Adder
}

func NewAdderMapper(adder Adder) *AdderMapper {
	return &AdderMapper{
		adder: adder,
		dict:  make(map[interface{}]Adder),
	}
}

func (m *AdderMapper) Feed(val interface{}) Adder {
	adder, ok := m.dict[val]
	if !ok {
		m.dict[val] = m.adder.Clone()
		adder = m.dict[val]
		m.adder.Incr()
	}

	return adder
}

func (m *AdderMapper) Foreach(do func(Adder, interface{})) {
	for val, adder := range m.dict {
		do(adder, val)
	}
}

type IntAdder int

var _ Adder = (*IntAdder)(nil)

func (i *IntAdder) Incr() {
	*i += 1
}

func (i *IntAdder) Clone() Adder {
	cpy := *i
	return &cpy
}
