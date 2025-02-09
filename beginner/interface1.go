package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak(){
	fmt.Printf("My name is %s and lastname is %s and age is %d\n", p.first, p.last, p.age)
}

func (sa secretAgent) speak(){
	fmt.Printf("Hi, I am a secret Agent. My name is %s and i have a license to kill %v", sa.first, sa.ltk)
}

type human interface{
	speak()
}

func saySomething(h human){
	h.speak()
}

func main(){
	p1 := person{
		first: "Elon",
		last: "Musk",
		age: 50,
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 35,
		},
		ltk: true,
	}
	saySomething(p1)
	saySomething(sa1)

}