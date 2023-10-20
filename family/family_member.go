package family

import . "fmt"

type Member struct {
	name string
	sex  string
	age  int
	role string
}

func (member *Member) GetMember() {
	Println("Name:", member.name, ", Sex:", member.sex, ", Age:", member.age, ", Role:", member.role)
}
func (member *Member) SetMember(name string, sex string, age int, role string) {
	member.name = name
	member.sex = sex
	member.age = age
	member.role = role
}
