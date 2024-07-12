package person_test

import (
	"strconv"
	"testing"

	"github.com/nyanshak/golang-learn-fp/cmd/person"
)

func BenchmarkImmutableCreatePerson(b *testing.B) {
	for i := range b.N {
		name := "person-" + strconv.Itoa(i)
		person.ImmutableCreatePerson(name, uint(i))
	}
}

func BenchmarkMutableCreatePerson(b *testing.B) {
	for i := range b.N {
		name := "person-" + strconv.Itoa(i)
		person.MutableCreatePerson(name, uint(i))
	}
}
