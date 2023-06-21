package emp

type Employee struct {
	eid  int
	name string
	EID  int
	Name string
}

func New(eid int, name string) *Employee {
	return &Employee{eid: eid, name: name}
}
