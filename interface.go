type Stringer interface {
	String() string
}

type Greeter struct {
	name	string
}

func (g *Greeter) String() string {
	return name
}
