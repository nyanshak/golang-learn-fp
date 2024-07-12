package person

type Person struct {
	Name string
	Age  uint
}

func ImmutableCreatePerson(name string, age uint) Person {
	p := Person{}
	p = ImmutableSetName(p, name)
	p = ImmutableSetAge(p, age)
	return p
}

func ImmutableSetName(p Person, name string) Person {
	p.Name = name
	return p
}

func ImmutableSetAge(p Person, age uint) Person {
	p.Age = age
	return p
}

func MutableCreatePerson(name string, age uint) *Person {
	p := &Person{}
	SetName(p, name)
	SetAge(p, age)
	return p
}

func SetAge(p *Person, age uint) {
	p.Age = age
}

func SetName(p *Person, name string) {
	p.Name = name
}
