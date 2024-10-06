package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

//
type Cow struct {
	food       string
	locomotion string
	noise      string
}
type Bird struct {
	food       string
	locomotion string
	noise      string
}
type Snake struct {
	food       string
	locomotion string
	noise      string
}

// func / method
//
func (c Cow) Eat()     { fmt.Println(c.food) }
func (c Cow) Move()    { fmt.Println(c.locomotion) }
func (c Cow) Speak()   { fmt.Println(c.noise) }
func (b Bird) Eat()    { fmt.Println(b.food) }
func (b Bird) Move()   { fmt.Println(b.locomotion) }
func (b Bird) Speak()  { fmt.Println(b.noise) }
func (s Snake) Eat()   { fmt.Println(s.food) }
func (s Snake) Move()  { fmt.Println(s.locomotion) }
func (s Snake) Speak() { fmt.Println(s.noise) }

func main() {
	var command, param1, param2 string
	var animal, cow, bird, snake Animal
	var animals map[string]Animal
	cow = Cow{"grass", "walk", "moo"}
	bird = Bird{"worms", "fly", "peep"}
	snake = Snake{"mice", "slither", "hsss"}
	animals = map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Usage: <newanimal|query> <param1> <param2>")
	fmt.Println("to add an animal")
	fmt.Println("newanimal animal_name <animal_type>")
	fmt.Println("Or to query animal information")
	fmt.Println("query <animal_type> <information>")
	fmt.Println("with possible value for <animal_type> cow,bird,snake")
	fmt.Println("with possible value for <information> eat,move,speak")
	for {
		command, param1, param2 = "", "", ""
		fmt.Printf(">")
		in.Scan()
		response := strings.ToLower(in.Text())
		// check if there are 3 words
		words := strings.Fields(response)
		if len(words) == 3 {
			command = (strings.Split(response, " "))[0]
			param1 = (strings.Split(response, " "))[1]
			param2 = (strings.Split(response, " "))[2]
			switch command {
			case "newanimal":
				{
					_, ok := animals[param1]
					if ok {
						fmt.Println("animal name " + param1 + " already known")
					} else {
						switch param2 {
						case "cow":
							{
								animals[param1] = cow
								fmt.Println("Created it!")
							}
						case "bird":
							{
								animals[param1] = bird
								fmt.Println("Created it!")
							}
						case "snake":
							{
								animals[param1] = snake
								fmt.Println("Created it!")
							}
						default:
							fmt.Println("unknow animal type " + param2)
						}
					}
				}
			case "query":
				{
					_, ok := animals[param1]
					switch ok {
					case true:
						{
							animal = animals[param1]
						}
					case false:
						{
							fmt.Println("unknow animal " + param1)
						}
					}
					switch param2 {
					case "eat":
						animal.Eat()
					case "move":
						animal.Move()
					case "speak":
						animal.Speak()
					default:
						fmt.Println("unknow information " + param2)
					}
				}
			}
		} else {
			fmt.Println("invalid request, try again")

		}
	}
}
