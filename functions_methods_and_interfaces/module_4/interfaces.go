/*
Write a program which allows the user to create a set of animals and to get information
about those animals. Each animal has a name and can be either a cow, bird, or snake.
With each command, the user can either create a new animal of one of the three types,
or the user can request information about an animal that he/she has already created.
Each animal has a unique defined by the user. Note that the user can define animals
of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.

Animal   Food eaten    Locomotion method     Spoken sound
cow       grass           walk                 moo

bird      worms           fly                  peep

snake     mice            slither              hsss 

Your program should present the user with a prompt, “>”, to indicate that the user can
type a request. Your program should accept one command at a time from the user, print
out a response, and print out a new prompt on a new line. Your program should continue
in this loop forever. Every command from the user must be either a “newanimal” command
or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string
is “newanimal”. The second string is an arbitrary string which will be the name of the
new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
 Your program should process each newanimal command by creating the new animal and printing
 “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”.
The second string is the name of the animal. The third string is the name of the information
requested about the animal, either “eat”, “move”, or “speak”. Your program should process each
query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically,
the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no
arguments and return no values. The Eat() method should print the animal’s food, the Move()
method should print the animal’s locomotion, and the Speak() method should print the animal’s
spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define
methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal
interface. When the user creates an animal, create an object of the appropriate type. Your program
should call the appropriate method when the user issues a query command.
*/
package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food string
	locomotion string
	noise string
}

func (a Cow) Eat() {
	fmt.Printf("a Cow eats %s\n", a.food)
}
func (a Cow) Move() {
	fmt.Printf("Cows %s\n", a.locomotion)
}
func (a Cow) Speak() {
	fmt.Printf("a Cow speaks like %s\n", a.noise)
}

type Bird struct {
	food string
	locomotion string
	noise string
}

func (a Bird) Eat() {
	fmt.Printf("a Bird eats %s\n", a.food)
}
func (a Bird) Move() {
	fmt.Printf("Birds %s\n", a.locomotion)
}
func (a Bird) Speak() {
	fmt.Printf("a Bird speaks like %s\n", a.noise)
}

type Snake struct {
	food string
	locomotion string
	noise string
}

func (a Snake) Eat() {
	fmt.Printf("a Snake eats %s\n", a.food)
}
func (a Snake) Move() {
	fmt.Printf("Snakes %s\n", a.locomotion)
}
func (a Snake) Speak() {
	fmt.Printf("a Snake speaks like %s\n", a.noise)
}

func CreateAnimal(aType string) Animal {
	var newA Animal
	switch aType {
	case "cow":
		newA = Cow{"grass", "walk", "moo"}
	case "bird":
		newA = Bird{"worms", "fly", "peep"}
	case "snake":
		newA = Snake{"mice", "slither", "hsss"}
	default:
		fmt.Printf("Animal \"%s\" is not in the list\n", aType)
	}
	return newA
}

func main() {
	var c, a, t string
	cow := Cow{"grass", "walk", "moo"}
	bird := Bird{"worms", "fly", "peep"}
	snake := Snake{"mice", "slither", "hsss"}
	aMap := map[string]Animal {
		"cow": cow,
		"bird": bird,
		"snake": snake,
	}

	for ; true; {

		fmt.Printf("type a request > ")
		fmt.Scanf("%s %s %s", &c, &a, &t)
		
		if c == "query" {
			animal, ok := aMap[a]
			
			if !ok {
				fmt.Printf("Animal %s is not in the list\n", a)
				continue
			}
			
			switch t {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Printf("Action \"%s\" does not exists\n", t)
			}
		} else if c == "newanimal" {
			newA := CreateAnimal(t)

			if newA == nil {
				continue
			}

			fmt.Printf("Created it!\n")
		} else {
			fmt.Printf("Command \"%s\" does not exist\n", c)
		}
	}

}
