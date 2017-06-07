package basics

// A Person has a GetOlder method that adds a year to its age.
// The problem is that, a Person could lie about its age by directly setting the value of the Age property,
// bypassing the GetOlder setter.
// Any idea on how to fix it?

func NewPerson(age int) *Person {
	if age < 1 || age > 120 {
		age = 1
	}
	return &Person{Age: age}
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetOlder() {
	p.Age++
}
