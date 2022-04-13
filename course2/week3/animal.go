package main

import "fmt"

type Animal struct{
	food, locomotion, noise string
}

func (v Animal) Eat(){
	fmt.Println(v.food)
}

func (v Animal) Move(){
	fmt.Println(v.locomotion)
}

func (v Animal) Speak(){
	fmt.Println(v.noise)
}



func main(){

	m := map[string]Animal{
		"cow" : Animal{"grass", "walk", "moo"},
		"bird": Animal{"worms", "fly", "peep"},
		"snake": Animal{"rat", "slither", "hsss"},
	}

	for {

		fmt.Println(">")
		an := "0"
		ac := "0"
		fmt.Scan(&an, &ac)
		if ac == "eat"{
			m[an].Eat()
		} else if ac=="move"{
			m[an].Move()
		} else if ac=="speak"{
			m[an].Speak()
		}
	}
}