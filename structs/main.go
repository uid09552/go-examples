package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	zipcode int
	street  string
}

func main() {
	jef := person{firstName: "jef", lastName: "morgan"}
	jim := person{firstName: "jim", contact: contactInfo{street: "abc", zipcode: 1}}
	jim.print()
	fmt.Printf("%+v", jim)
	fmt.Println(jim)
	fmt.Println(jef)
	var max person
	max.firstName = "max"
	max.lastName = "ho"
	fmt.Println(max)
	fmt.Printf("%+v", max)
	var hans person
	fmt.Printf("%+v", hans)
	jim.updateName("updatedJimName")
	jim.print()

}
func (p *person) updateName(n string) {
	p.lastName = n
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
