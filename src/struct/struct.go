package main

import "fmt"

type person struct {
	name string
	age  int
}

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	Skills
	int
	specialty string
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	} else {
		return p2, p2.age - p1.age
	}
}

func main() {
	var tom person

	tom.name, tom.age = "Tom", 18
	bob := person{age: 25, name: "Bob"}
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bp_Older.name, bp_diff)

	jane := Student{Human: Human{"Jane", 36, 100}, specialty: "Biology"}

	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her specialty is ", jane.specialty)

	jane.Skills = []string{"anatomy", "brain"}

	fmt.Println("Her skills are ", jane.Skills)

	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int)
}
