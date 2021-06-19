package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

type contactInfo struct {
	email string
}

func main() {

	defaultPerson := person{}  // default value: string = "", int/float=0, bool= false
	fmt.Println(defaultPerson) // {  0 {}}

	alex := person{
		lastName:  "Za",
		firstName: "Alex",
		age:       1,
		contact: contactInfo{
			email: "a@aws.com",
		},
	}
	fmt.Println(alex) // {Alex Za 1 {a@aws.com}}

	siri := person{"Si", "Ri", 1, contactInfo{"a@apple.com"}}
	siri.print()

}

func (p person) print() {
	fmt.Printf("%+v", p) // eg: {firstName:Si lastName:Ri age:1 contact:{email:a@apple.com}}
}
