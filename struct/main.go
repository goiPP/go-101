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
	siri.updateNamePassByValue("Sii") // will not take effect siri.firstName still be Si
	siriPointer := &siri
	siriPointer.updateNamePassByRef("Sii") // now siri.firstName will be Sii
}

func (p person) print() {
	fmt.Printf("%+v", p) // eg: {firstName:Si lastName:Ri age:1 contact:{email:a@apple.com}}
}

func (p person) updateNamePassByValue(newFirtName string) {
	p.firstName = newFirtName // this will not modify existing person, it passed by value (copy all props and paste into new container with new pointer)
	// and that temp person with new pointer will be use in this method
}

func (pointerToPerson *person) updateNamePassByRef(newFirtName string) {
	(*pointerToPerson).firstName = newFirtName // mean that we want to modify the value that pointerToPerson point/ref to
}
